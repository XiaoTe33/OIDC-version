var storage = window.localStorage;
const ul = document.querySelector('.details');
window.addEventListener('load',async()=>{

    //console.log(storage.getItem("token"));
    let url = `http://39.101.72.18:9090/user/info`
    const res = await fetch(url, {
      method : 'get',
      headers : {"Token": storage.getItem("token")}
    })
     let data = await res.json();
     console.log(data);
     if( data.status == 200){
        let url = `http://39.101.72.18:9090/order?status=${1}`
        const res = await fetch(url, {
          method : 'get',
          headers : {"Token": storage.getItem("token")}
        })
         let data = await res.json();
         console.log(data);
     }
     

})

setInterval(async()=>{
    console.log('111')
    const fd = new FormData()
    fd.append("token",storage.getItem("token"))
    const res = await fetch('http://39.101.72.18:9090/user/login/token', {
        method : 'post',
        body: fd,
    })
    let data = await res.json();
    console.log(data);
    if(data.status == 200){
        storage.setItem("token",data.token)
    }
},1000)



