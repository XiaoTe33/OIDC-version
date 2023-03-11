//alert("注册协议[审慎阅读]您在申请注册流程中点击同意前，应当认真阅读以下协议。请您务必审慎阅读、充分理解协议中相关条款内容,其中包括:1.与您约定免除或限制责任的条款:2.与您约定法律适用和管辖的条款:3.其他以粗体下划线标识的重要条款。如您对协议有任何疑问，可向平台客服咨询。[特别提示]当您按照注册页面提示填写信息阅读并同意协议且完成全部注册程序后，即表示您已充分阅读、理解并接受协议的全部内容。阅读协议的过程中,如果您不同意相关协议或其中任何条款约定,您应立即停止注册程序。《苏宁会员章程》《易付宝协议》《苏宁隐私政策》《易付宝隐私政策》《苏宁联盟在线协议》");


const psd = document.querySelector('#psw');
const username = document.querySelector('#username')
const repsd = document.querySelector('#repsw')
const phone = document.querySelector('#phone');
const email = document.querySelector('#email');
const btn = document.querySelector('#submit');
btn.addEventListener('click',async()=>{
  const fd = new FormData();
  fd.append("username",username.value);
  fd.append("password",psd.value);
  fd.append("rePassword",repsd.value);
  fd.append("phone",phone.value);
  fd.append("email",email.value);
  const res = await fetch('http://39.101.72.18:9090/user/register', {
      method: "post",
      body: fd,
  })
const data = await res.json();
console.log(data);
if(data.status != 200){
  alert(data.msg);
}else{
  let fd1 = new FormData();
  fd1.append("username",username.value);
  fd1.append("password",psd.value);
  const res = await fetch('http://39.101.72.18:9090/user/login', {
      method: "post",
      body: fd,
    })
    console.log(res.json());
    var storage = window.localStorage;
    storage.setItem("token",data.token);
    
}
})