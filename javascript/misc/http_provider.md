function provider_http( url, callback, params, type, headers )
{
    url = url || null;
    type = type || 'GET';
    params = params || {};
    headers = headers || {};
    callback = callback || function(){};

    headers['X-CSRF-TOKEN'] = 'ALGUN-TOKEN-PARA-AGREGAR-POR-DEFAULT';


    if( type == 'DELETE' || type == 'PUT' )//PARA EL REST
    {
        params['_method'] = type;
        type = 'POST';
    }
    //console.log("HTTP SENDING===>", 'type=',type, 'headers=', headers,'url=',  url,'params=', params)
    $.ajax
    ({
        type: type,
        headers: headers,
        url: url,
        data: params,
        success: function(data, textStatus, jQxhr) 
        {
            GLOBAL_response = data;
            callback();
        },
        error: function(jqXhr, textStatus, errorThrown) 
        {
            alert("ocurrio un error");
            console.log(jqXhr, textStatus, errorThrown);
        }
    });
}