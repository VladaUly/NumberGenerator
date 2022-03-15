var i = 0;
var txt;
var speed = 50;

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

var numberName = numberResult();
var chanelName = chanelResult();
if (numberName > 0 && numberName != "" && chanelName > 0 && chanelName !=""){
// Создаём новый объект XMLHttpRequest
var xhr = new XMLHttpRequest();

var params = 'number=' + encodeURIComponent(numberName)+ '&chanel=' + encodeURIComponent(chanelName)

// Конфигурируем его: GET-запрос на URL 
xhr.open('GET', '/generator?' + params, true);

// Отсылаем запрос
xhr.send();
xhr.onload = function() {
    let responseObj = xhr.response;
    console.log(responseObj)
    parseNum(responseObj)
  };
  
  xhr.onerror = function() { 
    alert(`Ошибка соединения`);
  };}
}

function parseNum(nums){
  document.getElementById("output").innerHTML = "OUTPUT : ";
  i= 0;
  txt = "[" + " " + nums + "]"
  typeWriter()

}

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


function typeWriter() {
  if (i < txt.length) {
    document.getElementById("output").innerHTML += txt.charAt(i);
    i++;
    setTimeout(typeWriter, speed / txt.length * 30);
  }
}