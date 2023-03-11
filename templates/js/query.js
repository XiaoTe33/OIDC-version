var queryText = document.querySelector('#query-box');
const btn = document.querySelector('#search');
const p1 = document.querySelector('#rangeField');
const p2 = document.querySelector('#method');
let mode;
for(let i = 0; i<p1.children.length;i++){
    if(p1.children[i].selected == true){
        mode = p1.children[i].value;
    }
}
for(let i = 0; i<p2.children.length;i++){
    if(p2.children[i].selected == true){
        mode = mode+p2.children[i].value;
    }
}

const name = document.querySelector('#queryField1');
const kind = document.querySelector('#queryField2');

let pictures = document.getElementsByClassName('goods-picture')
let goodsBoxes = document.getElementsByClassName('goods-box')
let goodsname =document.getElementsByClassName('name')
let price =document.getElementsByClassName('price')
let link =document.querySelectorAll('.goods-link')

btn.addEventListener('click',async()=>{
    let url;
    if(name.checked == true){
        url = `http://39.101.72.18:9090/goods?name=${queryText.innerHTML}&mode=${mode}`
    }else if( kind.checked == true){
        url = `http://39.101.72.18:9090/goods?kind=${queryText.innerHTML}&mode=${mode}`
    }
    const res = await fetch(url, {
        method: "get",
      })
      const data = await res.json();
      console.log(data);
      if(data.data.status ==200){
            for (let i = 0; i < 10; i++) {
                errMsg.style.visibility='hidden'
                if (i < data.data.length) {
                    goodsBoxes[i].style.display = 'block'
                    goodsname[i].innerHTML=data.data.name
                    // console.log(goodsBoxes[i].children[0]); 
                    // goodsBoxes[i].children[0].addEventListener('click',()=>{
                    //     location.assign('http://xiaote33.top:9090/goods?'+''+'='+''+'&mode='+'')
                    // })
                    price[i].innerHTML='￥'+data.data.price+'.00';
                    pictures[i].src = data.data.pictureLink;                    
                } else {
                    goodsBoxes[i].style.display = 'none'
                }
                if (i >= 9) {
                    break;
                }         
            }
      }
})
