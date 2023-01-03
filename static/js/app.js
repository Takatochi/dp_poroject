'use strict';
import {Activebtn,RouterHrefbtnactive} from "./pkg/active.js"

import {Getscript} from "./pkg/Ajax/responsescipt.js"
import {ObserverAppDOOM} from "./pkg/observer.js";


( function () {

    const appDom=document.getElementById("app")

    window.addEventListener("load", () => {
        const loaderInner = document.querySelector(".loader_inner");
        const loader = document.querySelector(".loader");

        setTimeout(() => {
            loaderInner.style.display = "none";
            loader.style.display = "none";

        }, 600);
    });


    RouterHrefbtnactive('barmenu','a')
   
   const itemGroup= Activebtn('barmenu',"a")
    const init=()=> {
        let router = new Router([
            new Route('New', 'listServer.html'),
            new Route('Dashboard', 'Dashboard.html'),
            new Route('Server', 'Server.html')
        ]);

    }
    init();

    const observe= new ObserverAppDOOM(appDom)

    const mutations = [
        {
            type: 'characterData',
            target: appDom,
            value: 'new value'
        },  {
            type: 'childList',
            target: appDom,
            addedNodes: ['new child node'],
            removedNodes: ['removed child node']
        }
    ];
    observe.onMutation(mutations)




})();


// function ObserverAppDOOMs (appDom){
//     let mutationObserver =   new MutationObserver(function(mutations) {
//
//         let event =  new CustomEvent("appDom", {
//             bubbles: true,detail:{observe:mutations}
//
//         })
//         console.log(event)
//         setTimeout(() =>  appDom.dispatchEvent(event));
//         $('#SettingHubModal').modal('hide')
//         $('#CreatePopModalCenter').modal('hide')
//     });
//
//     mutationObserver.observe(appDom, {
//         characterData: true,
//         childList: true,
//         subtree: false,
//         attributeOldValue: true,
//         characterDataOldValue: true,
//
//     });
//     console.log(mutationObserver)
// }