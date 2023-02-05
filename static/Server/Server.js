"use strict";


import {sendFile, ServerActivity} from "../js/datainterface/list.js";
import {appendTable, play, stop} from "./active.js";

(()=>{
    const app=document.getElementById("app");
    app.addEventListener('appDom',(e)=>{
        e.detail.observe.forEach(ob=> {
            if (!window.onloading&&window.location.hash.substr(1) === "Server") {
                if (location.ServerData === undefined)
                {
                    location.href = "/#"
                    return ;
                }
                main()
            }
        });
    });
})()


const main = ()=>{

    const serverData= new Map(
        [
            [
                0,{
            "ServerName":location.ServerData[0],
                "ServerPort":location.ServerData[1]
            },

            ]
        ]

    );
    infoCart(serverData.get(0).ServerName,serverData.get(0).ServerPort);

    play(serverData)
    stop(serverData)
    infoLoader(serverData.get(0).ServerPort)

    styleUse(serverData)


}

const infoCart = (name,port) => {
    const serveName= document.getElementById("ServerName");
    const servePort= document.getElementById("ServerPort");
    serveName.innerText=name;
    servePort.innerText=`:${port}`;
}
const infoLoader = (port)=>{
    let pastValue = document.querySelector(".js-value").innerHTML;
    $('input[type="file"]').change(function(){
        const file = document.getElementById('file-input').files[0];
        let value = $("input[type='file']").val();
        $('.js-value').text(value);
        sendFile(file,port).then(e=>{
            console.log(e)
        }).finally(()=>{
            setTimeout(()=>{
                $('.js-value').text(pastValue);
            },3000)
            appendTable(port)
        })
    });
}

const styleUse =(serverData)=>{
    const loader =document.querySelector('.file_loader')
    if (ServerActivity.has(serverData.get(0).ServerName)){
        loader.style.display='block';
    }
}
