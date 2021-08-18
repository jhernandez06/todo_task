import "expose-loader?exposes=$,jQuery!jquery";
import "bootstrap/dist/js/bootstrap.bundle.js";
import "@fortawesome/fontawesome-free/js/all.js";


$(() => {
   
    //let Url = $(location).attr('href');
    //let incomplete = "http://localhost:3000/?check_complet=false";
    //let completed = "http://localhost:3000/?check_complet=true"
    //let taskAll = "http://localhost:3000/"
    
   

    $("body").on("click",'.delete' ,function() {
        let deleteBtn = '<a href="/delete/'+$(this).attr('data-id')+'" data-method="DELETE" ><button type="button" class="btn btn-primary">Delete</button></a>';
        document.getElementById('deleteTask').innerHTML = deleteBtn;
     });

     //animation
      $('.check').on('mouseenter',aumentarText);
      $('.check').on('mouseleave',disminuirText);

     function aumentarText(){
       
         $(this).animate({ fontSize: '17px'},100)
     };
     function disminuirText(){

         $(this).animate({ fontSize: '16px'}, 100)
     };

});