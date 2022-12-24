
import {Activebtn,RouterHrefbtnactive} from "./pkg/active.js"
import {Getscript} from "./pkg/Ajax/responsescipt.js"



( ()=> {
    const appDom=document.getElementById("app")


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




    let mutationObserver = new MutationObserver(function(mutations) {

        let event = new CustomEvent("appDom", {
            bubbles: true,detail:{observe:mutations}
        })
        setTimeout(() => appDom.dispatchEvent(event));
        $('#SettingHubModal').modal('hide')
        $('#CreatePopModalCenter').modal('hide')
    });

    mutationObserver.observe(appDom, {
        characterData: true,
        childList: true,
        subtree: false,
        attributeOldValue: true,
        characterDataOldValue: true,

    });
    // console.log(window.loadspipt)
    //     loadscript(el.getAttribute('href'),"#New",urlnew)
    //new URLSearchParams({
    //         'user': 'user2'
    //     }))
    //{
    //         username: 'username',
    //         password: 'password'
    //     }, {
    //         headers: {
    //             'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
    //         }}

})();

