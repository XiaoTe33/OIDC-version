var storage = window.localStorage;

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
        storage.setItem("token",data.token);
    }
},2000000)

const list = document.querySelectorAll('.list-input');
const bg = document.querySelector('#bg');
const add = document.querySelector('.add');
const submit = document.querySelector('.addSubmit');
const addAddress = document.querySelector('#alert');
const close = document.querySelector('#close');
close.addEventListener('click',()=>{
    addAddress.classList.remove('active');
    bg.classList.remove('active');
})

add.addEventListener('click',()=>{
    addAddress.classList.add('active');
    bg.classList.add('active');
})

submit.addEventListener('click',async()=>{
    const fd = new FormData();
    fd.append("province",list[0].value);
    fd.append("city",list[1].value);
    fd.append("county",list[2].value);
    fd.append("detail",list[3].value);
    fd.append("tag",list[4].value);
      const res = await fetch('http://39.101.72.18:9090/user/address', {
          method: "post",
          headers : {"Token": storage.getItem("token")},
          body: fd,
      })
    let data = await res.json();
    console.log(data);
    if( data.status == 200){
    console.log('222');
  }
})