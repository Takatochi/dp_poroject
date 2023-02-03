import {ServerActivity, startServer, stopServer} from "../js/datainterface/list.js";


const appendAlert = (message, type) => {
    const alertPlaceholder = document.getElementById('Server')
    const wrapper = document.createElement('div')
    wrapper.innerHTML = [
        ` <div class = "alert alert-${type} d-flex align-items-center alert-dismissible fade show" role = "alert" >`,
        ` <svg class = "bi flex-shrink-0 me-2" width = "24" height = "24" role = "img" aria-label="${type}:" > `,
        `<use xlink:href = "#exclamation-triangle-fill" /> </svg>`,
        `<div>${message}</div>`,
        `<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`,
        `</div>`
    ].join('')

    alertPlaceholder.append(wrapper)
}
const play=(dataMap)=>{

    const loader =document.querySelector('.file_loader'),
    idPlay= document.getElementById('play'),
    idStop  = document.getElementById('stop')


    idPlay.addEventListener('click',()=> {
        startServer(dataMap).then(
            data => {
                appendAlert(data.data.message,"success")
                 loader.style.display='block';
                console.log(data)
            }
        ).catch(error=>{
            appendAlert(error.response.data.error,"danger")
        }).finally(_=>{
            ServerActivity.set(`${dataMap.get(0).ServerName}`,true)
            // disabled.
        })
    })

}
const stop=(dataMap)=>{
    const loader =document.querySelector('.file_loader'),
        idPlay= document.getElementById('play'),
    idStop  = document.getElementById('stop')

    idStop.addEventListener('click',()=> {
        stopServer(dataMap.get(0).ServerPort).then(
            data => {
                appendAlert(data.data.message,"success")
                loader.style.display='none';
            }
        ).catch(error=>{
            appendAlert(error.response.data.error,"warning")
        }).finally(_=>{
            ServerActivity.delete(`${dataMap.get(0).ServerName}`)
        })
    })


}
export {play,stop}