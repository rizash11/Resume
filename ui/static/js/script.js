document.addEventListener("DOMContentLoaded", function () {
  init();
});

function init() {
  // *** CALCULATING AGE ***
  const currentDate = new Date();
  const currentYear = currentDate.getFullYear();
  const currentMonth = currentDate.getUTCMonth() + 1;
  const currentDay = currentDate.getDate();
  let age = 25;

  if (currentMonth > 10 || (currentMonth === 10 && currentDay >= 29)) {
    age = currentYear - 1998;
  } else {
    age = currentYear - 1998 - 1;
  }

  const ageBlock = document.querySelector("#age");
  ageBlock.innerHTML = age.toString() + " years, born 29 October 1998";
  // ***********************
  //
  //
  //
  //
  //
  // *** SELECTING A PAGE ***
  if (document.getElementById("RadioPage1").checked) {
    switchPage(1);
  } else {
    switchPage(2);
  }
  // ************************
  //
  //
  //
  //
  //
  // *** SELECTING A SECTION ***
  switchSection();

  // ***************************
}

//
//
//
//
//
//

function switchPage(pageNum) {
  if (pageNum === 1) {
    document.querySelector("#page1").style.display = "block";
    document.querySelector("#page2").style.display = "none";
  } else if (pageNum === 2) {
    document.querySelector("#page1").style.display = "none";
    document.querySelector("#page2").style.display = "flex";
  }
}

function switchSection() {
  let sectionID = document.querySelector(
    'input[name="menuOption"]:checked'
  ).value;

  let sections = document.querySelectorAll("#page2 section");
  sections.forEach((section) => (section.style.display = "none"));
  document.querySelector("#page2 #" + sectionID).style.display = "block";
  document.querySelector(
    "#page2 #" + sectionID + "Label"
  ).style.backgroundColor = "rgb(230, 205, 185)";
}
