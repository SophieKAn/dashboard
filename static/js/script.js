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
					lab_title.innerText = l.title;
					cs_lab.appendChild(lab_title);
					widget.appendChild(cs_lab);
          for (i = l.start; i <= l.end; i++) {
						createLabMachine(cs_lab, i);
          }
        });
      }
    }
  }
  req.open('GET', url, true);
  req.send();
}

function createLabMachine(cs_lab, i) {
	let lab_machine = document.createElement('div');
	lab_machine.className = "lab_machine";
	lab_machine.innerText = i;
	cs_lab.appendChild(lab_machine);
}
