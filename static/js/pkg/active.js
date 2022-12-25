import {addList, SettingPupAdd} from "./ListUpdate.js";
import {initServer, listServer,deleteServer} from "../datainterface/list.js";

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

const ListBtnactive = (getElementById ,querySelectorAll)=>{
    const listGroup = document.getElementById(getElementById),
     itemGroup = listGroup.querySelectorAll(querySelectorAll)

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

                    }
                )

            }
        )


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
const Deletebtn=(getElementById,ElementById)=>{
    const listGroup = document.getElementById(getElementById),
        delete_btn = document.getElementById(ElementById)

    delete_btn.addEventListener('click', ()=>{

        
        const elementToDelete = listGroup.querySelector('[active="true"]');
        // Перевірити чи існує такий елемент
        if (!elementToDelete) {return}

        const key=elementToDelete.getAttribute("key")

            // Видалити елемент з документу
        elementToDelete.parentNode.removeChild(elementToDelete);

        const id = listServer[key].id
        deleteServer(id).then(_ => {
            listServer.splice(key, 1);
            changeKey(listGroup)
        })
            .catch(error => {
                console.log(error);
            });

    })







}
const changeKey=(listGroup)=>{

   const item=listGroup.querySelectorAll("button")
    item.forEach((e,index)=>{
        e.setAttribute('key',index-1 );
    })

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


        initServer(input.value)
            .then(data=>{
                Server.id=data.data.id
                Server.port=data.data.port
                Server.name=input.value
                listServer.push(Server)
                addList(Server, listServer.length-1)
        })
            .catch(error=>{
            console.log(error)
        })



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
function isFormFilled(form) {
    const inputs = form.querySelectorAll('input');
    for (let i = 0; i < inputs.length; i++) {
        if (!inputs[i].value) {
            return false;
        }
    }
    return true;
}
export {Activebtn,ListBtnactive,Createbtn,RouterHrefbtnactive,SettingHubModal,ServerBtn,Deletebtn};