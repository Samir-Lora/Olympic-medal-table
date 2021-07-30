$(document).ready(function () {
    $.ajax({
        method: "GET",
        url : '/api/countries',
        dataType: "json",
        success: function(data){
            $.each(data, function (i, item) { 
                var row = "<tr>"+
                "<td class='fw-bold text-center'>"+item.position+"</td>"+
                "<td class'fw-bold'>"+item.country+"</td>"+
                "<td class='fw-bold text-center'>"+item.goldmedalds+"</td>"+
                "<td class='fw-bold text-center'>"+item.silvermedalds+"</td>"+
                "<td class='fw-bold text-center'>"+item.bronzemedalds+"</td>"+
                "<td class='fw-bold text-center'>"+JSON.parse( item.goldmedalds+item.silvermedalds+item.bronzemedalds) +"</td>"
                $("#table>tbody").append(row)
            });
        }
    })
});
