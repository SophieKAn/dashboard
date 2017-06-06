package static

const Style = `
.widget,.cs_lab {
  display: -webkit-flex;
  -webkit-flex-flow: row wrap;
  display: flex;
  flex-flow: row wrap;
  align-items:stretch;

}

.widget {
  padding:15px;
  justify-content:space-around;
  align-content:stretch;
}


.cs_lab {
  background:#070707;
  width:200px;
  padding:20px;
  margin:10px;
  justify-content:flex-start;
  align-content:space-between;
}



.lab_title {
  width: 180px;
  height: 30px;
  text-align:center;
  color:white;
  font-size: 1.5em;
}


.lab_machine {
  margin:5px;
  background:gray;
  width:30px;
  height:30px;

  text-align:center;
  color:white;
  line-height:30px;
}
`
