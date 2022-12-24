import {addList, SettingPupAdd} from "./ListUpdate.js";
import {initServer, listServer} from "../datainterface/list.js";

const RouterHrefbtnactive =(ElementById,querySelectorAll)=>{
    const listGroup = document.getElementById('barmenu'),
        itemGroup = listGroup.querySelectorAll('a')
    itemGroup.forEach(btn=>{
        let href= btn.getAttribute('href')

        if (href===location.hash)
        {
            btn.classList.add('active')
        }
    })
}

const ListBtnactive = (getElementById ,querySelectorAll,funcElementById)=>{
    const listGroup = document.getElementById(getElementById)
    const itemGroup = listGroup.querySelectorAll(querySelectorAll)
    if (!itemGroup) {return}
        itemGroup.forEach(
            element => {
                element.addEventListener('click',
                    function () {
                        itemGroup.forEach(
                            e => {
                                e.classList.remove('active')
                                e.removeAttribute("active")
                            }
                        )
                        this.classList.add('active')
                        this.setAttribute("active", true);
                        Deletebtn(this, funcElementById)

                    }
                )

            }
        );

}
const Activebtn = (ElementById ,querySelectorAll)=>{

    const listGroup = document.getElementById(ElementById),
        itemGroup = listGroup.querySelectorAll(querySelectorAll)
    itemGroup.forEach(
        element => {
            element.addEventListener('click',
                function () {
                    itemGroup.forEach(
                        e => e.classList.remove('active')
                    )
                    this.classList.add('active')
                }
            )

        }
    );
    return itemGroup;
}
const Deletebtn=(elemtn,ElementById)=>{
    const status=elemtn.getAttribute("active")
    if (status){
        const element = document.getElementById(ElementById)
        element.addEventListener('click', ()=>{
            elemtn.remove()
            elemtn=null
        })
    }

}
const Createbtn=(ElementById,funcElementById)=>{
    let Server = new Object();
    const listGroup = document.getElementById(ElementById)
    const el=document.getElementById(funcElementById)
    el.addEventListener('click',()=>{
        $('#CreatePopModalCenter').modal('show')

    })
    const saveModal=document.getElementById("modal-btn-save")
    saveModal.addEventListener('click',()=>{
        $('#CreatePopModalCenter').modal('hide')
        const input= document.getElementById('inputServer')
        Server.id=0
        Server.name=input.value
        Server.port=2323

        initServer(input.value)
            .then(data=>{
                console.log(data)
        })
            .catch(error=>{
            console.log(error)
        })

        listServer.push(Server)
        console.log(listServer[listServer.length - 1])
        addList(Server, listServer.length-1)

    })
}
const SettingHubModal=(ElementById,obj)=>{

    const listGroup = document.getElementById(ElementById),
    itemGroup=listGroup.querySelectorAll('button')
     itemGroup.forEach(item=>  {
        if (item.getAttribute('active')) {
            item.addEventListener('dblclick', function () {
                $('#SettingHubModal').modal('show')
                SettingPupAdd(obj, this)
            })
        }

    })

}
const ServerBtn=(hash)=>{
    const btnServer=document.getElementById("btnServer")
    btnServer.addEventListener('click',()=>{
        location.hash=hash

    })
}
export {Activebtn,ListBtnactive,Createbtn,RouterHrefbtnactive,SettingHubModal,ServerBtn};