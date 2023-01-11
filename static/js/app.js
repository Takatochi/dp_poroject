'use strict';
import {Activebtn,RouterHrefbtnactive} from "./pkg/active.js"

import {Getscript} from "./pkg/Ajax/responsescipt.js"
import {ObserverAppDOOM} from "./pkg/observer.js";
import {Router} from "./pkg/router.js";

const init=()=> {
    console.log("1")
    let router = new Router([
        new Route('/', 'home.html',true),
        new Route('New', 'listServer.html'),
        new Route('Dashboard', 'Dashboard.html'),
        new Route('Server', 'Server.html')
    ]);
    console.log(router)
}
( function () {

    const appDom=document.getElementById("app")
    const loaderInner = document.querySelector(".loader_inner");
    const loader = document.querySelector(".loader");
    window.addEventListener("load", ()=>load(appDom,loaderInner,loader));

    RouterHrefbtnactive('barmenu','a')
   
   const itemGroup= Activebtn('barmenu',"a")





})();


const load = (appDom, loaderInner, loader) =>{
    setTimeout(() => {
        loaderInner.style.display = "none";
        loader.style.display = "none";
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
        init();

    }, 600);


}

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