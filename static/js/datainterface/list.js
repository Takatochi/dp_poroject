
import axios from "/static/js/pkg/axios.min.js"
const  listServer=[]

const getNewServer= async () => {
     await axios.post('/New')
        .then ((response)=> {
            listServer.push(...response.data)
        })
        .catch(function (error) {
            console.log(error);
        });


}


const initServer= async (name) => {

  return  await axios.post('/Server/init',
         {
             message:name
         },{
                headers:{
                    'Content-Type':'multipart/form-data; charset=UTF-8',
                }})

}
const deleteServer= async (id) => {
   return  axios.delete(`/Server/delete/server/${id}`)

}
export {listServer,getNewServer,initServer,deleteServer}