manageData();

function manageData() {
    $.ajax({
        dataType: 'json',
        url: "/camera",
    }).done(function (data) {
        manageRow(data.host,data.data);
    });
}

var keytmp=0;
/* Add new devices table row */
function manageRow(host,data) {
    var rows = '';
    $.each(data, function (key, value) {
        rows = rows + '<tr>';
        rows = rows + '<td><input type="text" class="form-control" value="' + key + '" readonly></td>';
        rows = rows + '<td><input type="text" class="form-control" value="' + value.nama + '"></td>';
        rows = rows + '<td><input type="text" class="form-control" value="' + value.rtsp + '"></td>';
        rows = rows + '<td><input type="text" class="form-control" value="' + value.path + '"></td>';
        rows = rows + '<td><small>' +host+ value.path + '.m3u8</small></td>';
        rows = rows + '</tr>';
        keytmp=key;
    });
    $("tbody").html(rows);
}


/* Create new devices */
$(".crud-submit").click(function (e) {
    e.preventDefault();
    var form_action = $("#create-devices").find("form").attr("action");
    var nama = $("#create-devices").find("input[name='nama']").val();
    var no_telp = $("#create-devices").find("input[name='no_telp']").val();
    $.ajax({
        dataType: 'json',
        type: 'POST',
        url: form_action,
        data: { nama: nama, no_telp: no_telp }
    }).done(function (data) {
        getPageData();
        $(".modal").modal('hide');
        toastr.success('devices Created Successfully.', 'Success Alert', { timeOut: 5000 });
    });


});
