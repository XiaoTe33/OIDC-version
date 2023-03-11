const details = document.querySelectorAll('.text');
const name = document.querySelector('#name')
var storage = window.localStorage;
const fileInp = document.querySelector('#fileInp');
const avator = document.querySelector('.touxiang');
const pic = document.querySelector('#avator');
const avatorbtn = document.querySelector('#avatorbtn');
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
       details[1].children[0].innerHTML = data.data.phone;
       details[2].children[0].innerHTML = data.data.email;
       name.innerHTML = data.data.username;
       pic.src = data.data.photo;
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


avator.onclick = function(){
  fileInp.click()
}
avatorbtn.addEventListener('click',async()=>{
    if(fileInp.files.length >0){
          const fd = new FormData();
        fd.append("photo",fileInp.files[0]);
        const res = await fetch('http://39.101.72.18:9090/user/photo', {
            method: "post",
            headers : {"Token": storage.getItem("token")},
            body: fd,
        })
      let data = await res.json();
      console.log(data);
      if( data.status == 200){
      console.log('修改成功，请刷新');
    }
  }
})

const push = document.querySelector('#push');
const del = document.querySelector('#delete');
const nikname = document.querySelector('#nicheng');

push.addEventListener('click',async()=>{
        const fd = new FormData();
        fd.append("username",details[0].children[0].value);
        fd.append("email",details[2].children[0].innerHTML);
        fd.append("phone",details[1].children[0].innerHTML);
        console.log(fd.get("email"));
        const res = await fetch('http://39.101.72.18:9090/user/info', {
            method: "put",
            headers : {"Token": storage.getItem("token")},
            body: fd,
        })
      //let data = await res.json();
      console.log(res);
      if(res.status == 200){
        alert('修改成功');
      }
})
 