var storage = window.localStorage;
const person = document.querySelector('#person');
const city = document.querySelector('#city');
const detailAddress = document.querySelector('#detailaddress');

const pic = document.querySelector('.pic');
const name = document.querySelector('.name');
const version = document.querySelector('.version');
const price = document.querySelector('.price');
const num = document.querySelector('#num');

window.addEventListener('load',async()=>{
const url = `http://39.101.72.18:9090/goods?goodsId=${storage.getItem("id")}`
    const res = await fetch(url, {
       method:"get"
           })
     const goodsdata = await res.json();
     console.log(goodsdata);
     console.log(goodsdata.data.pictureLink)
     if(goodsdata.status == 200){
        // storage.setItem('shopId',goodsdata.data.shopId);
        pic.children[0].src = goodsdata.data.pictureLink;
        name.innerHTML = goodsdata.data.name;
        price.innerHTML = '￥'+ goodsdata.data.price +'.00';
     }
//登录
    const resp = await fetch('http://39.101.72.18:9090/user/info', {
      method : 'get',
      headers : {"Token": storage.getItem("token")}
    })
     let persondata = await resp.json();
     console.log(persondata);
     person.innerHTML = persondata.data.username;
//地址 ok
    const responce = await fetch('http://39.101.72.18:9090/user/address', {
    method : 'get',
    headers : {"Token": storage.getItem("token")}
  })
   let data = await responce.json();
   console.log(data);
   if(data.status == 200){
    city.innerHTML = data.data[0].province + data.data[0].city;
    detailAddress.innerHTML = data.data[0].county + data.data[0].detail;
   }
})
//刷新token，保持登录
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


//提交订单
const buy = document.querySelector('.buy');
buy.addEventListener('click',async()=>{
    const fd = new FormData();
    fd.append("goodsId",storage.getItem("id"));
    fd.append("amount",Number(num.value));
    fd.append("style",1);
    fd.append("address",city.innerHTML+detailAddress.innerHTML);  
    const res = await fetch('http://39.101.72.18:9090/order', {
        method: "post",
        headers : {"Token": storage.getItem("token")},
        body: fd,
      })
      const data = await res.json();
      console.log(data);
    //   if(data.data.status == 200){
        
    //   }
})


//数量加减ok
const del1 = document.querySelector('.del1');
const add1 = document.querySelector('.add1');

del1.addEventListener('click',()=>{ 
    if(num.value > 1){
        num.value = num.value - 1;
    }
    if(num.value <= 1){
        console.log('111');
            //del1[i].style.cursor = 'not-allowed';
        del1.style.color = '#dedede';
    }
})
add1.addEventListener('click',()=>{
    num.value = Number(num.value)+ 1;
    if(num.value > 1){
        console.log('111')
        //del1[i].style.cursor = 'point';
        del1.style.color = '#666666';
    }
})

//const 