function initializePage(url) {
  let req = new XMLHttpRequest();
  req.onreadystatechange = () => {
    if (req.readyState === XMLHttpRequest.DONE) {
      if (req.status === 200) {
        let result = JSON.parse(req.responseText);
        let body = document.getElementsByTagName('body')[0];
        let widget = document.createElement('ul');
        widget.className = "widget";
        body.appendChild(widget);
        result.forEach(l => {
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
				updater();
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



function updater() {
	var conn = new WebSocket("ws://localhost:8080/upd");
	conn.onclose = function(evt) {
		console.log("Connection closed");
	}
	conn.onmessage = function(evt) {
		let machineData = JSON.parse(evt.data);
		console.log("updating");
		changeStatus(machineData);
	}
}


function changeStatus(machineData) {
  machineData.forEach(m => {
		let el = document.getElementById(m.hostname);
			if (m.status == 0) {
				el.style.background = "MediumSeaGreen";
			} else if (m.status == 1) {
				el.style.background = "MediumSlateBlue";
			} else {
				el.style.background = "gray";
			}
	});
}

function pad(n) {
	// http://stackoverflow.com/a/8089938/6279238
	// only for positive integers
	return (n < 10) ? ("0" + n.toString()) : String(n);
}
