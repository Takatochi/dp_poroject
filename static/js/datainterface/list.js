"use strict";
import axios from "/static/js/pkg/axios.min.js"
const  listServer=[]

const getNewServer= async () => {
      await axios
        .post('/New')
        .then ((response)=> {
            if (response.data.length > 0) {
                listServer.push(...response.data);
            }
        })
        .catch((error) => {
            console.log(error);
        });
}


const initServer=async (name) => {
  return  await axios
      .post('/Server/init',
         {
             message:name
         },{
                headers:{
                    'Content-Type':'multipart/form-data; charset=UTF-8',
                }})

}
const deleteServer= (id) => {
   return  axios.delete(`/Server/delete/server/${id}`)

}
const stopServer= (port) => {
    return  axios.delete(`/Server/Close/${port}`)

}
const startServer =async (serverMap)=>{
  return  await axios.post('/Server/Start',{

        port: serverMap.get(0).ServerPort,
        name: serverMap.get(0).ServerName,
    })
}

const sendFile = async (file)=>{

    const formData = new FormData();
    formData.append('file', file);

 return  await axios.post('/File/sql', formData)
        // .then(response => {
        //     console.log(response.data);
        // })
        // .catch(error => {
        //     console.log(error);
        // });
}
export {listServer,getNewServer,initServer,deleteServer,startServer,stopServer, sendFile}