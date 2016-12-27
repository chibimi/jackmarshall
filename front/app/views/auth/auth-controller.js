'use strict';

app.controller('AuthCtrl', ["$localStorage", "$http", function($localStorage, $http) {
    var scope = this;
    this.username = "";
    this.password = "";

    this.login = function(){
        $http.post(auth_endpoint + '/login', { login: scope.username, password: scope.password })
        .success(function (response) {
            if (response) {
                $localStorage.currentUser = { username: scope.login, token: response };
                $http.defaults.headers.common.Authorization = 'Bearer ' + response;
            }
        });
    }
}]);
