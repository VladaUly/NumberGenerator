let arrNum = []

function numberResult() { 
    var numberInp = document.getElementById("numberInp").value;
    var numberInt = parseInt(numberInp)
    if (numberInt > 0 && numberInp != ""){
        console.log("The number value=" + numberInt);
        return numberInt}
    else{
         alert("Please enter the integer value less than or equal to 0..");}
 
}

function chanelResult() { 
    var chanelInp = document.getElementById("chanelInp").value;
    var chanelInt = parseInt(chanelInp)
    if (chanelInt > 0 && chanelInp != ""){
        console.log("The chanel value=" + chanelInt);
        return chanelInt}
    else{
        alert("Please enter the integer value less than or equal to 0..");}
}

document.addEventListener("DOMContentLoaded", ready);

function ready() {
    animation();
    var button = document.getElementById("mainBtn")
    button.addEventListener("click",parseNumbers);
}

function parseNumbers() {
  arrNum = []
  var numberName = numberResult();
  var chanelName = chanelResult();
  if (numberName > 0 && numberName != "" && chanelName > 0 && chanelName !=""){
    var params = 'number=' + encodeURIComponent(numberName)+ '&chanel=' + encodeURIComponent(chanelName)
    let socket = new WebSocket("ws://127.0.0.1:8000/ws?" + params);
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = function(event) {
      console.log("Data is Recieved");
      let message = event.data;
      parseNum(message)
    }
    
    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!")
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
  }
}


function parseNum(nums){
  arrNum.push(nums)
  console.log(arrNum)

  document.getElementById("output").innerHTML = "[ OUTPUT : " + arrNum + " ]";}


function animation(){
    document.getElementById("rowImage").animate([
        // keyframes
        { transform: 'translate3D(0, 10px, 0)' },
        { transform: 'translate3D(0, -10px, 0)' }
      ], {
        // timing options
        duration: 900,
        iterations: Infinity,
        direction: "alternate"
      })

      
}


