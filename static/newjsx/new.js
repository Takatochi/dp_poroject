"use strict";
import _, {map, times} from "../js/pkg/underscore-esm.js";
import {CreateServer, ListBtnactive, SettingHubModal, ServerBtn, Deletebtn} from "/static/js/pkg/active.js"
import {listServer,getNewServer} from "../js/datainterface/list.js";
import {addList} from "../js/pkg/ListUpdate.js";
// import {ObserverAppDOOM} from "../js/pkg/observer";
const load=()=>{
    setTimeout(() => {
        loaderInner.style.display = "none";
        loader.style.display = "none";
    }, 600);
}
(()=>{
            getNewServer().then().catch(error=>{
            console.log(error+" 12")
        }).finally(
            LoadNews
        )
})()

function LoadNews(){

    const app=document.getElementById("app")

    app.addEventListener('appDom',(e)=>{

            if(!window.onloading&&window.location.hash.substr(1)==="New") {

                listServer.forEach((obj, index) => {
                    addList(obj, index)
                })
                ListBtnactive('list-group', "button")
                CreateServer("list-group", "create_btn")
                ServerBtn("#Server")
                Deletebtn('list-group', "delete_btn")
                SettingHubModal('list-group',listServer)
                // app.addEventListener('click',()=>{
                //
                // })

            }


    })
}








//пошук в масиві (на майбутне)
// function contains(arr, elem) {
//     return arr.find((i) => i === elem) != -1;
// }
// console.log(contains(listServer,"Names"))
