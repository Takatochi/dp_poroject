

const  Getscript=(urlscript)=>{
    window.loadspipt="add"
    urlscript.forEach(urls=>{
        axios.get(urls[0])
            .then(function (response) {

                // handle success
                const app = document.getElementById('app')
                const s= document.createElement('script')

                s.type=urls[1];
                s.src=urls[0];
                app.append(s)

            })
            .catch(function (error) {
                // handle error
                console.log(error);
            })
            .finally(function () {
                // always executed
            });
    })

    // Usage

}

export {Getscript}