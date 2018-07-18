Vue.component('cliente', {

    template: `
    <div>
        <div class="card">
            <h5 class="card-header">CLIENTE</h5>
            <div class="card-body">

                <div class="form-group">
                    <label> Razon Social</label>
                    <input type="text" class="form-control" v-model="cliente.razon_social">
                </div>
                <div class="form-group">
                    <label> Direccion</label>
                    <input type="text" class="form-control" v-model="cliente.direccion">
                </div>
                <div class="form-group">
                    <label> RUC</label>
                    <input type="text" class="form-control" v-model="cliente.ruc">
                </div>
                <div class="form-group">
                    <label> Telefono</label>
                    <input type="text" class="form-control" v-model="cliente.telefono">
                </div>
                <div class="form-group pull-right">
                    <input type="button" class="btn btn-primary" value="GUARDAR" v-on:click="add()">
                </div>
                
            </div>
        </div>
        <br>
        <div class="table-responsive">
            <table class="table table-striped table-bordered table-hover">
                <thead>
                    <tr>
                        <th>RAZON SOCIAL</th>
                        <th>DIRECCION</th>
                        <th>RUC</th>
                        <th>TELEFONO</th>
                        <th>ACCIONES</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="c in clientes">
                        <td v-text="c.razon_social"></td>
                        <td v-text="c.direccion"></td>
                        <td v-text="c.ruc"></td>
                        <td v-text="c.telefono"></td>
                        <td></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>`,

    data: function()
    {
        return {
            cliente:{},
            clientes: []
        }
    },

    methods:
    {
        add: function()
        {
            this.clientes.push( this.cliente );
            this.cliente = {};
        }
    },
    
    computed: function()
    {

    }
    

})