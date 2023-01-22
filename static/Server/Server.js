"use strict";


import {startServer} from "../js/datainterface/list.js";
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


}

const infoCart = (name,port) => {
    const serveName= document.getElementById("ServerName");
    const servePort= document.getElementById("ServerPort");
    serveName.innerText=name;
    servePort.innerText=`:${port}`;
}