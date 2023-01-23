"use strict";


import {sendFile} from "../js/datainterface/list.js";
import {play,stop} from "./active.js";

(()=>{
    const app=document.getElementById("app");
    app.addEventListener('appDom',(e)=>{
        e.detail.observe.forEach(ob=> {
            if (!window.onloading&&window.location.hash.substr(1) === "Server") {
                if (location.ServerData === undefined)
                {
                    location.href = "/#"
                    return 0
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
    stop(serverData.get(0).ServerPort)
    infoLoader()

}

const infoCart = (name,port) => {
    const serveName= document.getElementById("ServerName");
    const servePort= document.getElementById("ServerPort");
    serveName.innerText=name;
    servePort.innerText=`:${port}`;
}
const infoLoader = ()=>{
    let pastValue = document.querySelector(".js-value").innerHTML;
    $('input[type="file"]').change(function(){
        const file = document.getElementById('file-input').files[0];
        let value = $("input[type='file']").val();
        $('.js-value').text(value);
        sendFile(file).then(e=>{
            console.log(e)
        }).finally(()=>{
            setTimeout(()=>{
                $('.js-value').text(pastValue);
            },3000)
        })
    });
}