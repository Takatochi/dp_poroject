import _,{map} from "../js/pkg/underscore-esm.js";
import {Createbtn, ListBtnactive,SettingHubModal,ServerBtn} from "/static/js/pkg/active.js"
import {listServer,getNewServer} from "../js/datainterface/list.js";
import {addList} from "../js/pkg/ListUpdate.js";
(()=>{
    getNewServer().then(
        LoadNews
    )
})()

function LoadNews(){
    const app=document.getElementById("app")

    app.addEventListener('appDom',(e)=>{
        e.detail.observe.forEach(ob=>{

            if(ob.addedNodes[0].tagName!=="SCRIPT"&&window.location.hash.substr(1)==="New") {

                listServer.forEach((obj, index) => {
                    addList(obj, index)
                })

                Createbtn("list-group", "create_btn")
                ServerBtn("#Server")
                ListBtnactive('list-group', "button", "delete_btn")
                app.addEventListener('click',()=>{
                   SettingHubModal('list-group',listServer)

                })

            }

        })
    })
}








//пошук в масиві (на майбутне)
// function contains(arr, elem) {
//     return arr.find((i) => i === elem) != -1;
// }
// console.log(contains(listServer,"Names"))
