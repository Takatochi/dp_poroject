import {startServer, stopServer} from "../js/datainterface/list.js";

const play=(dataMap)=>{
    const idPlay= document.getElementById('play')
    idPlay.addEventListener('click',()=> {
        startServer(dataMap).then(
            data => {
                console.log(data)
            }
        )
    })


}
const stop=(port)=>{
    const idPlay= document.getElementById('stop')
    idPlay.addEventListener('click',()=> {
        stopServer(port).then(
            data => {
                console.log(data)
            }
        )
    })


}
export {play,stop}