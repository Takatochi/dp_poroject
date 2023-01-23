'use strict';
import axios from "/static/js/pkg/axios.min.js"
function Router(routes) {
    try {
        if (!routes) {
            throw 'error: routes param is mandatory';
        }
        this.constructor(routes);
        this.init();
    } catch (e) {
        console.error(e);   
    }
}
Router.prototype.handleNoHashRoute = function(route) {
    if (!route.hash) {
        this.goToRoute(route.htmlName);
    }
}
Router.prototype = {
    routes: undefined,
    rootElem: undefined,
    constructor: function (routes) {
        this.routes = routes;
        this.rootElem = document.getElementById('app');
    },
    init: function () {
        let r = this.routes;
        (function(scope, r) { 
            window.addEventListener('hashchange', function (e) {
                scope.hasChanged(scope, r);
            });
        })(this, r);
        this.hasChanged(this, r);
    },
    hasChanged: function(scope, r){
        if (window.location.hash.length > 0) {
            for (let i = 0, length = r.length; i < length; i++) {
                let route = r[i];
                if(route.isActiveRoute(window.location.hash.substr(1))) {
                    scope.goToRoute(route.htmlName);
                }
            }
        } else {
            for (let i = 0, length = r.length; i < length; i++) {
                let route = r[i];
                if(route.default) {
                    scope.goToRoute(route.htmlName);
                }
                // Call the new function to handle routes without a hash

            }
        }
    },
    goToRoute: function(htmlName) {
        window.onloading=true;
        const template = `<div class="loaderApp" style="display: block;">
        <div class="loaderApp-inner ball-grid-pulse" >
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
        </div>
    </div>`
        this.rootElem.innerHTML = template;

        axios.get(`/static/${htmlName}`)
            .then(response => {
                setTimeout(()=>{
                    this.rootElem.innerHTML = response.data;
                    window.onloading=false;
                 },700)

            })
    }
};
export {Router}