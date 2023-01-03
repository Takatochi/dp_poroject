

const IDServer=()=>{
   return window.IdServer
}
const redirectionURL=()=>{
    if(IDServer()!==undefined)
      return
    location.href="/"

}
((id)=>{

    const app=document.getElementById("app")

    app.addEventListener('appDom',(e)=>{

        e.detail.observe.forEach(ob=> {

            if (window.location.hash.substr(1) === "Server") {

                redirectionURL()
                console.log(id())

            }
        })
    })
})(IDServer)
