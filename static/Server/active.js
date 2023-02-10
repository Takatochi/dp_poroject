import {ServerActivity, shower, startServer, stopServer} from "../js/datainterface/list.js";


const appendAlert = (message, type) => {
    const alertPlaceholder = document.getElementById('Server')
    const wrapper = document.createElement('div')
    wrapper.innerHTML = [
        ` <div class = "alert alert-${type} d-flex align-items-center alert-dismissible fade show" role = "alert" >`,
        ` <svg class = "bi flex-shrink-0 me-2" width = "24" height = "24" role = "img" aria-label="${type}:" > `,
        `<use xlink:href = "#exclamation-triangle-fill" /> </svg>`,
        `<div>${message}</div>`,
        `<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`,
        `</div>`
    ].join('')

    alertPlaceholder.append(wrapper)
}

const appendTable = async (port)=>{


    const VarType =[]
    await shower(port).then(data=>{
        if (data.data.type != null) {
            VarType.push(...data.data.type)
            return;
        }

    }).catch(error=>{
        appendAlert(error.response.data.error,"warning")
    })

    if (VarType.length<0) return;

    VarType.map((typesVar,index)=>{

        const alertPlaceholder = document.getElementById('accordionFlushType')
        const wrapper = document.createElement('div')
        let Naming=[]
        let SourceType=[]
        for (const property in typesVar.var.Columns) {
            console.log(`${property}: ${typesVar.var.Columns[property]}`);
            Naming+=`<th>${property}</th>`
            SourceType+=`<td>${typesVar.var.Columns[property]}</td>`

        }
        wrapper.innerHTML = [
           ` <div class="accordion-item">`,
                `<h2 class="accordion-header" id="flush-heading${index}">`,
                   `<button id="collapsed_btn"class="accordion-button collapsed" type="button" data-bs-toggle="collapse",
                            data-bs-target="#flush-collapse${index}" aria-expanded="false" aria-controls="flush-collapse${index}">`,
                        `${typesVar.tableName}`,
                    `</button>`,
                `</h2>`,
            `<div id="flush-collapse${index}" class="accordion-collapse collapse" aria-labelledby="flush-heading${index}"
                     data-bs-parent="#accordionFlushType">`,
            `<div class="table-responsive">`,
            `<table class="table">`,
            `<thead>`,
            `<tr>`,
            `<th>#</th>`,
             Naming,
            `</tr>`,
            `</thead>`,
            `<tbody>`,
            `<tr>`,
            `<th scope="row">1</th>`,
            SourceType,
            `</tr>`,
            `</tbody>`,
            ` </table>`,
            `</div>`,
            `</div>`,
            `</div>`,
        ].join('')

        alertPlaceholder.append(wrapper)
    })

}


const play=  (dataMap)=>{

    const loader =document.querySelector('.file_loader'),
    idPlay= document.getElementById('play'),
    idStop  = document.getElementById('stop')


    idPlay.addEventListener('click',()=> {
        startServer(dataMap).then(
            data => {
                appendAlert(data.data.message,"success")
                 loader.style.display='block';
                console.log(data)
            }
        ).catch(error=>{
            appendAlert(error.response.data.error,"danger")
        }).finally(_=>{
            ServerActivity.set(`${dataMap.get(0).ServerName}`,true)

            // disabled.
        })
    })

}
const stop=(dataMap)=>{
    const loader =document.querySelector('.file_loader'),
        idPlay= document.getElementById('play'),
    idStop  = document.getElementById('stop')

    idStop.addEventListener('click',()=> {
        stopServer(dataMap.get(0).ServerPort).then(
            data => {
                appendAlert(data.data.message,"success")
                loader.style.display='none';
            }
        ).catch(error=>{
            appendAlert(error.response.data.error,"warning")
        }).finally(_=>{
            ServerActivity.delete(`${dataMap.get(0).ServerName}`)
        })
    })


}
export {play,stop,appendTable}