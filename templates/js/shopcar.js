var storage = window.localStorage;
let details = document.querySelector('.details');
const del = document.querySelector('.delete');
let yes = document.querySelector('.yes');
let totalPrice = document.querySelector('.totalprice');
let totalNum = document.querySelector('.totalnum');
let address;

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
        const res = await fetch('http://39.101.72.18:9090/user/car', {
            method : 'get',
            headers : {"Token": storage.getItem("token")}
          })
           let data = await res.json();
           console.log(data);
           if(data.status == 200){
            for(let i = 0; i<data.data.length; i++){
                let ul = document.createElement('ul');
                ul.classList.add('flex');
                details.appendChild(ul);
                for( let j=1;j<=8;j++){
                    let li = document.createElement('li');
                    var classname = `list${j}`;
                    li.classList.add(classname);
                    ul.appendChild(li);
                    if(j==6){
                        li.classList.add('flex');
                        li.classList.add('num');
                        let btn1 = document.createElement('button');
                        li.appendChild(btn1);
                        btn1.classList.add('del1');
                        btn1.classList.add('btn');
                        btn1.innerHTML = '-'
                        let input = document.createElement('input');
                        input.classList.add('num1');
                        input.setAttribute("type","text")
                        input.value = 1;
                        li.appendChild(input);
                        let btn2 = document.createElement('button');
                        btn2.classList.add('add1');
                        btn2.classList.add('btn');
                        btn2.innerHTML = '+'
                        li.appendChild(btn2);
                        //在生成的同时绑定事件监听
                        btn1.addEventListener('click',()=>{ 
                                if(input.value > 1){
                                    input.value = input.value - 1;
                                }
                                if(input.value <= 1){
                                    console.log('111');
                                    //del1[i].style.cursor = 'not-allowed';
                                    btn1.style.color = '#dedede';
                                }
                            })
                        btn2.addEventListener('click',()=>{
                            input.value = Number(input.value)+ 1;
                            if(input.value > 1){
                                console.log('111')
                                //del1[i].style.cursor = 'point';
                                btn1.style.color = '#666666';
                            }
                        })
                        

                    }
                    if( j==1){
                        let input = document.createElement('input');
                        input.setAttribute("type","checkbox");
                        input.classList.add('check');
                        li.appendChild(input);
                        del.addEventListener('click',async()=>{
                                if( input.checked == true){
                                    let url = `http://39.101.72.18:9090/user/car/${Number(i)}`
                                    const res = await fetch(url,{
                                        method:"delete",
                                        headers : {"Token": storage.getItem("token")},
                                    })
                                    const data = await res.json();
                                    console.log(data);
                                    //console.log(url);
                                }
                        })
                        let m = 0;
                        yes.addEventListener('click',function(){
                            m++;
                            if (  m % 2 != 0 ){                               
                                    input.checked = true;
                            }else{            
                                    input.checked = false;
                            }   
                        });
                        // if( input.checked = true){
                        //     console.log('999')
                        //     totalNum.innerHTML = Number(totalNum.innerHTML) + Number(num[i]);
                        //     totalPrice.innerHTML = Number(totalPrice.innerHTML) + Number(sum[i]);
                        // }
                    }
                    if(j==2){
                        li.classList.add('pic');
                        let img = document.createElement('img');
                        li.appendChild(img);
                        img.src = data.data[i].pictureLink
                    }
                    if(j==3){
                        li.classList.add('name');
                        li.innerHTML = data.data[i].name;
                    }
                    if(j==4){
                        li.classList.add('version');
                    }
                    if(j==5){
                        li.classList.add('price');
                        li.innerHTML = Number(data.data[i].price);
                    }
                    if(j==7){li.classList.add('sum')}
                    if(j==8){
                        li.innerHTML = '移入关注'
                    }
                }
            }
           }
     }
    const responce = await fetch('http://39.101.72.18:9090/user/address', {
    method : 'get',
    headers : {"Token": storage.getItem("token")}
  })
   let addressdata = await responce.json();
   console.log(addressdata);
   if(data.status == 200){
    address = addressdata.data[0].province + addressdata.data[0].city + addressdata.data[0].county + addressdata.data[0].detail;
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

console.log(details.children[1])
//更新总价和总数 no
for(let i = 0 ;i<details.children.length;i++){  
    if( details.children[i].children[0].checked == true){
        console.log('999');
        totalNum.innerHTML = Number(totalNum.innerHTML) + Number(details.children[i].children[4].value);
        totalPrice.innerHTML = Number(totalPrice.innerHTML) + Number(details.children[i].children[3].innerHTML)*Number(details.children[i].children[4].value);
    }
}


//去结算
const btn = document.querySelector('.totalbtn');
btn.addEventListener('mouseenter',()=>{
    for(let i = 0; i<details.children.length;i++){
        if(details.children[i].children[0].checked == true){
            btn.style.backgroundColor = '#ff8000';
        }
    }
})
btn.addEventListener('click',async()=>{
    for(let i = 0; i<details.children.length;i++){
        if( details.children[i].children[0].checked == true){
          storage.setItem("id",)
          //通过购物车的id去查找对应商品id，再传过去
        }
    }
})
