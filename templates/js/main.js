

const phonenum = document.querySelector('.phonenum');
const example = document.querySelectorAll('.more');
const header = document.querySelector('.header2');
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
       phonenum.innerText = data.data.phone;
       phonenum.classList.remove('none');
       header.children[0].classList.add('none');
       header.children[1].classList.add('none');
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
},2000000)

let num = 0;
let circle = 0;
  // 小圆点的点击事件
  const imgbox = document.querySelector('.images')
  const ul = document.querySelector('.picture');
  let dots = document.querySelector('.dots');
  const width = imgbox.offsetWidth;
  for (let i = 0; i < ul.children.length; i++) {
    let li =document.createElement('li');
    //li.setAttribute('index',i);
    dots.append(li);
    li.addEventListener('click',()=>{
      for( let j = 0; j < dots.children.length; j++){
        dots.children[j].classList.remove('active');
      }
      li.classList.add('active');
      num = i;
      circle =i;
      //移动ul
      const width = imgbox.offsetWidth;
      console.log(width);
      animate(ul, -i*width);
    })
}
dots.children[0].classList.add('active');

//克隆第一张图(在生成li之后克隆)
let first = ul.children[0].cloneNode(true);
ul.appendChild(first);
//左右按钮
const left = document.querySelector('.zuo');
const right = document.querySelector('.you');
// let num = 0;
// let circle = 0;
right.addEventListener('click',()=>{
  if(num == ul.children.length - 1){
    ul.style.left = 0;
    num=0;
  }
  num++;
  animate(ul, -num*width);
  circle++;
  if(circle == dots.children.length){
    circle=0;
  }
  for(let i = 0; i<dots.children.length;i++){
    dots.children[i].classList.remove('active');
  }
  dots.children[circle].classList.add('active');
})


left.addEventListener('click',()=>{
  if(num == 0){
    num = ul.children.length - 1;
    ul.style.left = -num*width + 'px';
  }
  num--;
  animate(ul, -num*width);
  circle--;
  if(circle < 0){
    circle = dots.children.length - 1;
  }
  for(let i = 0; i<dots.children.length;i++){
    dots.children[i].classList.remove('active');
  }
  dots.children[circle].classList.add('active');
})


const timer = setInterval(function () {
        right.click();
    }, 10000);
    
imgbox.addEventListener('mouseenter', ()=> {
        left.style.display = 'block';
        right.style.display = 'block';
    });
imgbox.addEventListener('mouseleave',()=> {
        left.style.display = 'none';
        right.style.display = 'none';
    });





// //注册事件

for( let i = 0; i < example.length; i++){
    example[i].id = '3055483356794';
    example[i].addEventListener('click',()=>{
        let storage = window.localStorage;
        storage.setItem('id',example[i].id)
    })
}

//退出登录
const signOut = document.querySelector('.sign-out');
signOut.addEventListener('click',async()=>{
  const res = await fetch('http://39.101.72.18:9090/user/info', {
    method : 'get',
    headers : {"Token": "123456789"}
  })
   let data = await res.json();
   console.log(data);
   if( data.status != 200){
    console.log('333')
     //phonenum.innerText = data.data.phone;
     phonenum.classList.add('none');
     header.children[0].classList.remove('none');
     header.children[1].classList.remove('none');
     storage.setItem("token","123456789");
   }
})
