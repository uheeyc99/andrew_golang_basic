

var andrew_xmlhttp;
function loadXMLDoc(url,cfunc){
    andrew_xmlhttp=new XMLHttpRequest();
    andrew_xmlhttp.onreadystatechange=cfunc;
    andrew_xmlhttp.open("GET",url,true);
    andrew_xmlhttp.send();
}
function AndrewLoadText(url,id){
    loadXMLDoc(url,function(){
        if (andrew_xmlhttp.readyState==4 && andrew_xmlhttp.status==200){
            document.getElementById(id).innerHTML=andrew_xmlhttp.responseText;
        }
    });

}


function andrewAlert(str)
{
    alert(str);
}

function andrewLazy()
{
    $('img.lazy').lazyload(
        {
            //th/eshold : 200 ,
            //placeholder : "img/PoweredByMacOSX.gif"
        }
    );
}
