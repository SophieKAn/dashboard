function initializePage(url) {
  let req = new XMLHttpRequest();
  req.onreadystatechange = () => {
    if (req.readyState === XMLHttpRequest.DONE) {
      if (req.status === 200) {
        let json = JSON.parse(req.responseText);
        let body = document.getElementsByTagName('body')[0];
        let widget = document.createElement('ul');
        widget.className = "widget";
        body.appendChild(widget);
        json["machineRanges"].forEach(l => {
          let cs_lab = document.createElement('ul');
          cs_lab.className = "cs_lab";
          let lab_title = document.createElement('header');
          lab_title.className = "lab_title";
          lab_title.innerText = l.prefix;
          cs_lab.appendChild(lab_title);
          widget.appendChild(cs_lab);
          for (i = l.start; i <= l.end; i++) {
            createLabMachine(cs_lab, i);
          }
        });
        updater(json["interface"], json["port"], json["machineIdentifiers"]);
      }
    }
  }
  req.open('GET', url, true);
  req.send();
}

function createLabMachine(cs_lab, i) {
  let lab_machine = document.createElement('div');
  lab_machine.className = "lab_machine";
  lab_machine.id = cs_lab.firstChild.innerText + "-" + pad(i) + ".***REMOVED***";
  lab_machine.innerText = i;
  cs_lab.appendChild(lab_machine);
}



function updater(interf, port, machineIdentifiers) {
  var conn = new WebSocket("ws://" + interf + ":" + port + "/upd");
  conn.onclose = function(evt) {
    console.log("Connection closed");
  }
  conn.onmessage = function(evt) {
    let machineData = JSON.parse(evt.data);
    console.log("updating");
    changeStatus(machineData, machineIdentifiers);
  }
}


function changeStatus(machineData, machineIdentifiers) {
  machineData.forEach(m => {
    let el = document.getElementById(m.hostname);
    machineIdentifiers.forEach(s => {
      if (m.status == s.name) {
        el.style.background = s.color;
      }
    });
	  var idArray = [];
	  for (var i in machineIdentifiers) {
      idArray.push(machineIdentifiers[i].name);
    }
    if (!isInArray(m.status, idArray)) {
			el.style.background = "gray";
		}
  });
}

function pad(n) {
  // http://stackoverflow.com/a/8089938/6279238
  // only for positive integers
  return (n < 10) ? ("0" + n.toString()) : String(n);
}

function isInArray(value, array) {
  return array.indexOf(value) > -1;
}
