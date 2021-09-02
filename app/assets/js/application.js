import "expose-loader?exposes=$,jQuery!jquery";
import "bootstrap/dist/js/bootstrap.bundle.js";
import "@fortawesome/fontawesome-free/js/all.js";


$(() => {
   
    $("body").on("click",'.delete' ,function() {
        let x = $(this).attr('data-id');
        //let deleteBtn = '<a href="/delete/'+$(this).attr('data-id')+'" data-method="DELETE" ><button type="button" class="btn btn-primary">Delete</button></a>';
        //let deleteBtn = `<%=linkTo(deleteTaskIDPath({ task_id: ${x} }), {class: "btn btn-danger", "data-method": "DELETE", "data-confirm": "Are you sure?", body: "Destroy"}) %>`;
       let deleteBtn = `<a href="/tasks/delete/${x}" data-method="DELETE" data-confirm="Are you sure?" class="btn btn-danger btn-xs" data-remote="true" data-type="script">Delete</a>`

        document.getElementById('deleteTask').innerHTML = deleteBtn;
     });
     $("body").on("click",'#btn-delete' ,function() {
        let deleteBtn = '<a href="/user/delete/'+$(this).attr('data-id')+'" data-method="DELETE" class="btn btn-danger">Delete</a>';
        document.getElementById('deleteuser').innerHTML = deleteBtn;
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
     $(function() {
        $('[data-toggle="tooltip"]').tooltip()
      })

});