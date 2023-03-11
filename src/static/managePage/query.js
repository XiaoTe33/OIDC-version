
var queryMode = '20'
var btn = document.getElementById('search')
var query = function () {


    //1.获取元素
    var pictures = document.getElementsByClassName('goods-picture')
    var goodsBoxes = document.getElementsByClassName('goods-box')
    var errMsg = document.getElementById('msg')
    var name =document.getElementsByClassName('name')
    var price =document.getElementsByClassName('price')
    var link =document.querySelectorAll('.goods-link')
    var queryField =document.querySelectorAll
    //2.获取数据
    //2.1获取查询参数
    var queryText = document.getElementById('query-box');
    //2.2发送请求
    ajax({
        method: 'get',
        url: 'http://www.xiaote33.top:9090/goods',
        data: {
            name: queryText.value,
            mode: queryMode,
        },
        //3.提取前十条,拼接到html上
        success: function (resp) {
            console.log(resp);
            var i = 0
            for (i = 0; i < 10; i++) {
                errMsg.style.visibility='hidden'
                if (i < resp['data'].length) {
                    goodsBoxes[i].style.display = 'block'
                    name[i].innerHTML=resp['data'][i]['name']
                   //if
                    console.log(goodsBoxes[i].children[0]); 
                    goodsBoxes[i].children[0].addEventListener('click',()=>{
                        location.assign('http://xiaote33.top:9090/goods?'+''+'='+''+'&mode='+'')
                    })
                
                    price[i].innerHTML='$'+resp['data'][i]['price']+'.00'
                    pictures[i].src = resp['data'][i]['pictureLink']
                    
                } else {
                    goodsBoxes[i].style.display = 'none'
                }

                if (i >= 9) {
                    break
                }         
            }

        },
        err: function (resp) {
            console.log(resp);
            for (var i = 0; i < 10; i++) {
                goodsBoxes[i].style.display = 'none'

            }
            // errMsg.innerHTML='s'
            errMsg.innerHTML=resp['msg']

            errMsg.style.visibility='visible'
        }
    })
}