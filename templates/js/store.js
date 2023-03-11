var storage = window.localStorage;

window.addEventListener('load',async()=>{
     let url = `http://39.101.72.18:9090/shop/queryAll?id=${storage.getItem("shopId")}`
    const res = await fetch(url, {
        method : 'post',
    })
    let data = await res.json();
    console.log(data);
})