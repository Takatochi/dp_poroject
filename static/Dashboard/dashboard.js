/* globals Chart:false, feather:false */
"use strict";


(()=>{

  const app=document.getElementById("app")
  app.addEventListener('appDom',(e)=>{

    e.detail.observe.forEach(ob=>{

      if(ob.addedNodes[0].tagName!=="SCRIPT"&&window.location.hash.substr(1)==="Dashboard") {
        (function () {

          feather.replace({ 'aria-hidden': 'true' })

          // Graphs
          let ctx = document.getElementById('myChart')
          // eslint-disable-next-line no-unused-vars
          let myChart = new Chart(ctx, {
            type: 'line',
            data: {
              labels: [
                'Sunday',
                'Monday',
                'Tuesday',
                'Wednesday',
                'Thursday',
                'Friday',
                'Saturday'
              ],
              datasets: [{
                data: [
                  5,
                  4,
                  5,
                  4,
                  3,
                  112,
                  1434
                ],
                lineTension: 0,
                backgroundColor: 'transparent',
                borderColor: '#007bff',
                borderWidth: 4,
                pointBackgroundColor: '#007bff'
              }]
            },
            options: {
              scales: {
                yAxes: [{
                  ticks: {
                    beginAtZero: false
                  }
                }]
              },
              legend: {
                display: false
              }
            }
          })
        })()

      }

    })
  })
})()


