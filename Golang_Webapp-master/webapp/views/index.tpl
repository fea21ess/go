<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>雳鉴插桩式IAST-BeegoWebapp测试用例集合</title>
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" media="screen">
    <script src="https://cdn.bootcss.com/angular.js/1.6.3/angular.min.js" charset="utf-8"></script>
    <style media="screen">
        thead tr td {
            background-color: #f1f1f1
        }
    </style>
</head>

<body>
<div ng-app="myapp" ng-controller="main">
    <div class="container" id="main">
        <div class="row">
            <div class="col-xs-12 col-sm-8 col-sm-offset-2">
                <h3 class="text-center">雳鉴插桩式IAST测试用例集合</h3>
                <br/>
                <table class="table table-striped">
                    <thead>
                    <tr>
                        <td>测试用例</td>
                        <td>用例路径</td>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $k,$v := .rows1}}
                    <tr>
                        <td>{{$v.name}}</td>
                        <td><a target="_blank" ng-href="{{$v.path}}">{{$v.path}}</a></td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<script type="text/javascript">
    var app = angular.module('myapp', []);

    app.controller('main', ['$scope', '$http',
		function($scope, $http) {
            $scope.testcases = []
        }
    ]);
</script>

</body>
</html>
