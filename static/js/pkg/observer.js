'use strict';
function ObserverAppDOOM (appDom) {
    this.appDom = appDom;

    this.mutationObserver = new MutationObserver(this.onMutation.bind(this)
    );

}

ObserverAppDOOM.prototype.observeOptions = {
    characterData: true,
    childList: true,
    subtree: false,
    attributeOldValue: true,
    characterDataOldValue: true
}

ObserverAppDOOM.prototype.onMutation = function (mutations) {


    this.mutationObserver.observe(this.appDom, this.observeOptions);
    this.hideModals();
    this.dispatchEvent(mutations);
}

ObserverAppDOOM.prototype.hideModals = function () {
    $('#SettingHubModal').modal('hide');
    $('#CreatePopModalCenter').modal('hide');
}

ObserverAppDOOM.prototype.dispatchEvent = function (mutations) {
    let event =  new CustomEvent("appDom", {
        bubbles: true, detail: {observe: mutations}
    });
    setTimeout(() =>  this.appDom.dispatchEvent(event),10);
    // this.appDom.dispatchEvent(event);
}
export  {ObserverAppDOOM}