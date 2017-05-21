'use strict';

var app = angular.module('jackmarshall', [
    'ngRoute',
    'ui.bootstrap',
    'ngAnimate',
    'ngDraggable',
    'ngStorage',
    'angular-uuid'
]);

app.config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
    $routeProvider.when('/auth/login', {
        templateUrl: 'views/auth/login.html',
    });
    $routeProvider.when('/auth/new', {
        templateUrl: 'views/auth/new-user.html',
    });
    $routeProvider.when('/tournament/list', {
        templateUrl: 'views/tournamentList/tournament-list.html',
    });
    $routeProvider.when('/tournament/:id', {
        templateUrl: '/views/tournamentDetails/tournament-details.html',
    });
    $routeProvider.otherwise({redirectTo: '/auth/login'});
}]);
