const pic = document.querySelector('#pic');
const bigpic = document.querySelector('#bigpic');
const mask = document.querySelector('#mask');
const headline = document.querySelector('#headline');
const address = document.querySelector('.address')
var storage = window.localStorage;
//读取localstorage，渲染页面
window.addEventListener('load',async()=>{
    
    const url = `http://39.101.72.18:9090/goods?goodsId=${storage.getItem("id")}`
    const res = await fetch(url, {
       method:"get"
           })
     const goodsdata = await res.json();
     console.log(goodsdata);
     console.log(goodsdata.data.pictureLink)
     if(goodsdata.status == 200){
        storage.setItem('shopId',goodsdata.data.shopId);
        pic.children[0].src = goodsdata.data.pictureLink;
        bigpic.children[0].src = goodsdata.data.pictureLink;
        headline.innerHTML = goodsdata.data.name;
     }
//登录
    const resp = await fetch('http://39.101.72.18:9090/user/info', {
      method : 'get',
      headers : {"Token": storage.getItem("token")}
    })
     let persondata = await resp.json();
     console.log(persondata);
//地址
    const responce = await fetch('http://39.101.72.18:9090/user/address', {
    method : 'get',
    headers : {"Token": storage.getItem("token")}
  })
   let data = await responce.json();
   console.log(data);
   if(data.status == 200){
    address.innerHTML = data.data[0].province + data.data[0].city + data.data[0].county;
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




//放大镜
pic.addEventListener('mouseenter',()=>{
    bigpic.style.display = 'block';
    mask.style.display = 'block';
})
pic.addEventListener('mouseleave',()=>{
    bigpic.style.display = 'none';
    mask.style.display = 'none';
})
pic.addEventListener('mousemove',(event)=>{
    let clientX=event.clientX;
    let clientY=event.clientY;
    let eleX=pic.getBoundingClientRect().left;
    let eleY=pic.getBoundingClientRect().top;

    let left=clientX-eleX-mask.offsetWidth/2;
    let top=clientY-eleY-mask.offsetHeight/2;

    if(left<=0){
        left=0;
    }else if(left>=pic.clientWidth-mask.offsetWidth){
        left=pic.clientWidth-mask.offsetWidth;
    }

    if(top<=0){
        top=0;
    }else if(top>pic.clientHeight-mask.offsetHeight){
        top=pic.clientHeight-mask.offsetHeight;
    }
    mask.style.left=left+'px';
    mask.style.top=top+'px';
    bigpic.scrollLeft=left*2;
    bigpic.scrollTop=top*2;
})


//收藏 ok
const yuan = document.querySelector('.yuan');
const shoucang = document.querySelector('.shoucang');
const collect = document.querySelector('.collect');
collect.addEventListener('mouseenter',()=>{
    yuan.classList.add('hover1');
    shoucang.classList.add('hover2');
})
collect.addEventListener('mouseleave',()=>{
    yuan.classList.remove('hover1');
    shoucang.classList.remove('hover2');
})
collect.addEventListener('click',async()=>{
    let url = `http://39.101.72.18:9090/user/star/${storage.getItem("id")}`
    const res = await fetch(url,{
        method:"put",
        headers : {"Token": storage.getItem("token")},
    })
    //let data = res.json();
    console.log(res);
    if(res.status == 200){
        yuan.style.backgroundColor = '#ff7000';
        shoucang.style.color = '#ff7000';
        shoucang.innerHTML = '已收藏';
    }
})


//加购数量 ok
const num = document.querySelector('#num')
const del = document.querySelector('.del1');
const add = document.querySelector('.add1');
const buy = document.querySelector('.buyitnow');
const addtocar = document.querySelectorAll('.addtocar');
del.addEventListener('click',()=>{ 
    if(num.value > 1){
        num.value = num.value - 1;
    }
    if(num.value <= 1){
        console.log('111');
        //del.style.cursor = 'not-allowed';
        del.style.color = '#dedede';
    }
})
add.addEventListener('click',()=>{
    num.value = Number(num.value)+ 1;
    if(num.value > 1){
        console.log('111')
        //del.style.cursor = 'point';
        del.style.color = '#666666';
    }
})


//立即购买，跳转订单页面
// buy.addEventListener('click',async()=>{
//     const fd = new FormData();
//     fd.append("goodsId",storage.getItem("id"));
//     fd.append("amount",Number(num.value));
//     fd.append("style",1);
//     fd.append("address",address.innerHTML);  
//     console.log(fd.get('amount'));
//     const res = await fetch('http://39.101.72.18:9090/order', {
//         method: "post",
//         headers : {"Token": storage.getItem("token")},
//         body: fd,
//       })
//       const data = await res.json();
//       console.log(data);
//       if(data.data.status == 200){
//         buy.children.href = './order.html'
//       }
// })


//加入购物车 ok
for(let i = 0; i<addtocar.length;i++){
addtocar[i].addEventListener('click',async()=>{
    let url = `http://39.101.72.18:9090/user/car/${storage.getItem("id")}`
    const res = await fetch(url, {
        method: "put",
        headers : {"Token": storage.getItem("token")}
      })
      const data = await res.json();
      console.log(data);
})
}


//选择款式




//获取评论 no
const comment = document.querySelector('.comment1');
comment.addEventListener('click',async()=>{
    console.log('333')
    console.log(storage.getItem("id"))
    //const url = `http://39.101.72.18:9090/evaluation?goodsId=${storage.getItem("id")}`
    const res = await fetch('http://39.101.72.18:9090/evaluation?goodsId=3052467651255', {
        method:"get",
       })
    const data = await res.json();
    console.log(data);
    // if(data.status == 200){
    //     //渲染评论框
    // }
})


//店铺
// const store = document.querySelector('.store');
// store.addEventListener('click',async()=>{

// })

