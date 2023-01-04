
const addList=  (obj,key)=>{
    const group= document.getElementById("list-group")
    const el = document.createElement("button");
    el.className="list-group-item list-group-item-action";
    el.type="button";
    el.innerHTML = `<a>${obj.name} <samp>Port:${obj.port}</samp></a> `;
    el.setAttribute("key",key)
    group.appendChild(el);

}
const SettingPupAdd=(obj,btn)=>{
   const key=btn.getAttribute('key')
    const serverData=document.getElementById("serverData")
   const Servername=serverData.querySelectorAll('a')
    window.IdServer=key
    const datalist=[]
    Object.keys(obj[key]).forEach((keys)=> {
        if(keys==='id') return
            datalist.push(obj[key][keys])
    })

    Servername.forEach((data,i)=>{
          data.innerText=datalist[i]

    })
}
export {addList,SettingPupAdd}