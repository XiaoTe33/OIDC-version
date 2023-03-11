var storage = window.localStorage;
let icon_rmb = document.querySelector('.icon-rmb');
let rmb = document.querySelector('.rmb');
let history = document.querySelector('.icon-history');
let qiandao = document.querySelector('.icon-xianxing-15');
let hongbao = document.querySelector('.icon-hongbao');
   

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


icon_rmb.addEventListener('mouseenter',()=>{
    console.log('111')
    // rmb.classList.add('moved')
})