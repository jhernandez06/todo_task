import "expose-loader?exposes=$,jQuery!jquery";
import "bootstrap/dist/js/bootstrap.bundle.js";
import "@fortawesome/fontawesome-free/js/all.js";


$(() => {
   
    let Url = $(location).attr('href');
    let incomplete = "http://localhost:3000/tasks?check_complet=false";
    let completed = "http://localhost:3000/tasks?check_complet=true"
    let taskAll = "http://localhost:3000/tasks"
    
    if (Url == incomplete || Url == taskAll){
        $('#add-task').removeClass('d-none')
    } 
    
    if (Url == completed){
        $('#incomplete').addClass('font-weight-lighter');
    }else if (Url == incomplete){
        $('#completed').addClass('font-weight-lighter');
    }else{
       
        $('#completed').addClass('font-weight-lighter');
        $('#incomplete').addClass('font-weight-lighter');
    };

     function  currentDate( ){
         
        var week = ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'];
        var month = ["January","February","March","April","May","June","July","August","September","October","November","December"]
        const d = new Date();
        let dayWeek = week[d.getDay()] ;
        let dayMonth = d.getDate();
        let Month =  month[d.getMonth() + 1];
        let year = d.getFullYear();

        document.getElementById("date").innerHTML = `${dayWeek} ${dayMonth} , ${Month} ${year}`;

    };
    currentDate();

    $("body").on("click",'#callDelete' ,function() {
        let r = '<a href="/tasks/delete/'+$(this).attr('data-id')+'" data-method="DELETE" ><button type="button" class="btn btn-primary">Delete</button></a>';
        document.getElementById('deleteTask').innerHTML = r;
     });

});