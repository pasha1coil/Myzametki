function show_popap(id_popap) {
    var id = "#" + id_popap;
    $(id).addClass('active');
    }
function close_popap(id_popap) {
var id = "#" + id_popap;
$(".overlay").removeClass("active");
}

function opentab(id){
    window.open('http://127.0.0.1:8000/show?id='+id);
}

function closetab(){
    window.close();
}

function delZamet(id){
    window.open('http://127.0.0.1:8000/del?id='+id);
    window.close();
}



