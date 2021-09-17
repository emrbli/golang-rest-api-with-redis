// main.js
console.log("hello world");

new Vue({
  el: "#app",
  
  data() {
    return {
      name: "Modanisa!",
      todos: [],
      inputTodo: ""
    };
  },
  
  mounted () {
    axios
      .get('http://localhost:8082/keys')
      .then(response => (this.todos = response.data.split("\n"))).catch(function (error) {
        console.log(error);
    }

    );
  },
  
  methods: {
    createTodo() {
      if (this.inputTodo != "") {
        this.todos.push(this.inputTodo);
        this.putData(this.inputTodo)
        this.inputTodo = "";
        
      }
    },

    deleteTodo(todo) {
      
       console.log(todo)
       this.deleteData(todo)
       this.todos.splice(this.todos.indexOf(todo), 1);
    },
    putData(contact){
        console.log(contact);
        var today = new Date()
        const date = today.getFullYear() + '-' + (today.getMonth()+1) + '-' + today.getDate();
         axios.put("http://localhost:8082/keys"+
        '?key='+contact+'&value='+date
        ).then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
        });
        
    },
    deleteData(memId) {
        axios
          .delete("http://localhost:8082/keys/" + memId);
      },
       
  }
});
