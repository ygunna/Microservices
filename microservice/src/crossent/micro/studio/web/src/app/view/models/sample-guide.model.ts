export class SampleGuide {
  frontApplication: string;
  frontCongroller: string;
  frontData: string;
  frontApplicationProperties: string;
  frontBootstrapProperties: string;
  frontPom: string;
  backApplication: string;
  backCongroller: string;
  backData: string;
  backApplicationProperties: string;
  backBootstrapProperties: string;
  backPom: string;
  indexHtml: string;
  albumsJs: string;
  appJs: string;
  errorsJs: string;
  albumsHtml: string;
  errorsHtml: string;
  frontManifest: string;
  backManifest: string;

  password: string ="${password}";
  spring_cloud_version: string = "${spring-cloud.version}";


  constructor() {
    this.frontApplication = `package com.crossent.microservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@EnableDiscoveryClient
@SpringBootApplication
public class FrontApplication {

	public static void main(String[] args) {
		SpringApplication.run(FrontApplication.class, args);
	}

}


@Configuration
@EnableSwagger2
class SwaggerConfig {
	@Bean
	public Docket api() {
		return new Docket(DocumentationType.SWAGGER_2)
				.select().apis(RequestHandlerSelectors.any())
				.paths(PathSelectors.ant("/api/**"))
				.build();
	}
}`;
    this.frontCongroller = `package com.crossent.microservice;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.client.RestTemplate;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

import javax.annotation.PostConstruct;
import java.net.URI;

@Controller
@RefreshScope
public class FrontController {

	private static final Logger logger = LoggerFactory.getLogger(FrontController.class);

	@Autowired
	private RestTemplate searchClient;


	@RequestMapping("/api/search")
	@ResponseBody
	public Data[] search(Model model) {
		System.out.println("request /api/search");

		URI uri = URI.create("http://apigateway/back/search/get");
		Data[] obj = searchClient.postForObject(uri, null, Data[].class);

		model.addAttribute("datas", obj);

		return obj;
	}

	@RequestMapping("/api/searchTwo")
	@ResponseBody
	public Data[] searchTwo(Model model) {
		System.out.println("request /api/searchTwo");

		URI uri = URI.create("http://apigateway/back/search/getTwo");
		Data[] obj = searchClient.postForObject(uri, null, Data[].class);

		model.addAttribute("datas", obj);

		return obj;
	}

	@RequestMapping("/api/searchThree")
	@ResponseBody
	public Data[] searchThree(Model model) {
		System.out.println("request /api/searchThree");

		URI uri = URI.create("http://apigateway/back/search/getThree");
		Data[] obj = searchClient.postForObject(uri, null, Data[].class);

		model.addAttribute("datas", obj);

		return obj;
	}


	@Value("${this.password}:")
	String password;

	@PostConstruct
	private void postConstruct() {

		System.out.println("My password is: " + password);
	}

}

@Configuration
class AppConfiguration {

	@LoadBalanced
	@Bean
	RestTemplate restTemplate() {
		return new RestTemplate();
	}
}

`;


    this.frontData = `package com.crossent.microservice;

public class Data {
    String id;
    String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }
}
`;
    this.frontApplicationProperties =
      "eureka.instance.hostname=${CF_INSTANCE_INTERNAL_IP} \n"+
      "eureka.instance.nonSecurePort=${PORT} \n";
    this.frontBootstrapProperties =
      "spring.application.name=front \n";
    this.frontPom = `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>

	<groupId>com.crossent.microservice</groupId>
	<artifactId>front</artifactId>
	<version>0.0.1-SNAPSHOT</version>
	<packaging>jar</packaging>

	<name>front</name>
	<description>Demo project for Spring Boot</description>

	<parent>
		<groupId>org.springframework.boot</groupId>
		<artifactId>spring-boot-starter-parent</artifactId>
		<version>1.5.9.RELEASE</version>
		<relativePath/> <!-- lookup parent from repository -->
	</parent>

	<properties>
		<project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
		<project.reporting.outputEncoding>UTF-8</project.reporting.outputEncoding>
		<java.version>1.8</java.version>
		<spring-cloud.version>Edgware.RELEASE</spring-cloud.version>
	</properties>

	<dependencies>
		<dependency>
			<groupId>org.springframework.cloud</groupId>
			<artifactId>spring-cloud-starter-config</artifactId>
		</dependency>
		<dependency>
			<groupId>org.springframework.cloud</groupId>
			<artifactId>spring-cloud-starter-eureka</artifactId>
		</dependency>
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-actuator</artifactId>
		</dependency>
    <dependency>
      <groupId>org.springframework.cloud</groupId>
      <artifactId>spring-cloud-starter-hystrix</artifactId>
    </dependency>		
		<dependency>
			<groupId>io.springfox</groupId>
			<artifactId>springfox-swagger2</artifactId>
			<version>2.3.1</version>
		</dependency>
		<dependency>
			<groupId>io.springfox</groupId>
			<artifactId>springfox-swagger-ui</artifactId>
			<version>2.3.1</version>
		</dependency>

		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-test</artifactId>
			<scope>test</scope>
		</dependency>
		
		<dependency>
			<groupId>io.pivotal.spring.cloud</groupId>
			<artifactId>spring-cloud-services-cloudfoundry-connector</artifactId>
			<version>1.6.1.RELEASE</version>
		</dependency>
		<dependency>
			<groupId>io.pivotal.spring.cloud</groupId>
			<artifactId>spring-cloud-services-spring-connector</artifactId>
			<version>1.6.1.RELEASE</version>
		</dependency>		
		
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<dependency>
			<groupId>org.webjars</groupId>
			<artifactId>jquery</artifactId>
			<version>2.2.4</version>
		</dependency>
		<dependency>
			<groupId>org.webjars</groupId>
			<artifactId>bootstrap</artifactId>
			<version>4.0.0</version>
		</dependency>
		<dependency>
			<groupId>org.webjars</groupId>
			<artifactId>angularjs</artifactId>
			<version>1.2.16</version>
		</dependency>
		<dependency>
			<groupId>org.webjars</groupId>
			<artifactId>angular-ui</artifactId>
			<version>0.4.0-2</version>
		</dependency>				
	</dependencies>

	<dependencyManagement>
		<dependencies>
			<dependency>
				<groupId>org.springframework.cloud</groupId>
				<artifactId>spring-cloud-dependencies</artifactId>
				<version>${this.spring_cloud_version}</version>
				<type>pom</type>
				<scope>import</scope>
			</dependency>
		</dependencies>
	</dependencyManagement>

	<build>
		<plugins>
			<plugin>
				<groupId>org.springframework.boot</groupId>
				<artifactId>spring-boot-maven-plugin</artifactId>
			</plugin>
		</plugins>
	</build>


</project>
`;
    this.backApplication = `package com.crossent.microservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@EnableDiscoveryClient
@SpringBootApplication
public class BackApplication {

	public static void main(String[] args) {
		SpringApplication.run(BackApplication.class, args);
	}
}

@Configuration
@EnableSwagger2
class SwaggerConfig {
	@Bean
	public Docket api() {
		Docket docket = new Docket(DocumentationType.SWAGGER_2)
				.select().apis(RequestHandlerSelectors.any())
				.paths(PathSelectors.ant("/search/**"))
				.build();
		docket.host("apigateway");
		docket.pathMapping("back");

		return docket;
	}
}
`;
    this.backCongroller = `package com.crossent.microservice;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import javax.annotation.PostConstruct;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/search")
@RefreshScope
public class BackController {

	private static final Logger logger = LoggerFactory.getLogger(BackController.class);


	@RequestMapping(value="/get", method = RequestMethod.POST)
	public List<Data> search(Model model) {
		System.out.println("request /serarc/get");

		Data data = new Data();
		data.setId("1");
		data.setName("name");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		return list;
	}

	@RequestMapping(value="/getTwo", method = RequestMethod.POST)
	public List<Data> searchTwo(Model model) {
		System.out.println("request /serarc/getTwo");

		Data data = new Data();
		data.setId("1");
		data.setName("name");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		data = new Data();
		data.setId("2");
		data.setName("name");

		list.add(data);

		return list;
	}

	@RequestMapping(value="/getThree", method = RequestMethod.POST)
	public List<Data> searchThree(Model model) {
		System.out.println("request /serarc/getThree");

		Data data = new Data();
		data.setId("1");
		data.setName("name");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		data = new Data();
		data.setId("2");
		data.setName("name");

		list.add(data);

		data = new Data();
		data.setId("3");
		data.setName("name");

		list.add(data);

		return list;
	}


}

`;
    this.backData = `package com.crossent.microservice;
import io.swagger.annotations.ApiModel;

@ApiModel(description="Data")
public class Data {
    String id;
    String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }
}
`;
    this.backApplicationProperties =
      "eureka.instance.hostname=${CF_INSTANCE_INTERNAL_IP} \n"+
      "eureka.instance.nonSecurePort=${PORT} \n";
    this.backBootstrapProperties =
      "spring.application.name=back \n";
    this.backPom = `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>

	<groupId>com.crossent.microservice</groupId>
	<artifactId>back</artifactId>
	<version>0.0.1-SNAPSHOT</version>
	<packaging>jar</packaging>

	<name>back</name>
	<description>Demo project for Spring Boot</description>

	<parent>
		<groupId>org.springframework.boot</groupId>
		<artifactId>spring-boot-starter-parent</artifactId>
		<version>1.5.9.RELEASE</version>
		<relativePath/> <!-- lookup parent from repository -->
	</parent>

	<properties>
		<project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
		<project.reporting.outputEncoding>UTF-8</project.reporting.outputEncoding>
		<java.version>1.8</java.version>
		<spring-cloud.version>Edgware.RELEASE</spring-cloud.version>
	</properties>

	<dependencies>
		<dependency>
			<groupId>org.springframework.cloud</groupId>
			<artifactId>spring-cloud-starter-config</artifactId>
		</dependency>
		<dependency>
			<groupId>org.springframework.cloud</groupId>
			<artifactId>spring-cloud-starter-eureka</artifactId>
		</dependency>	
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-actuator</artifactId>
		</dependency>
    <dependency>
      <groupId>org.springframework.cloud</groupId>
      <artifactId>spring-cloud-starter-hystrix</artifactId>
    </dependency>		
		<dependency>
			<groupId>io.springfox</groupId>
			<artifactId>springfox-swagger2</artifactId>
			<version>2.3.1</version>
		</dependency>
		<dependency>
			<groupId>io.springfox</groupId>
			<artifactId>springfox-swagger-ui</artifactId>
			<version>2.3.1</version>
		</dependency>

		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-test</artifactId>
			<scope>test</scope>
		</dependency>
		
		<dependency>
			<groupId>io.pivotal.spring.cloud</groupId>
			<artifactId>spring-cloud-services-cloudfoundry-connector</artifactId>
			<version>1.6.1.RELEASE</version>
		</dependency>
		<dependency>
			<groupId>io.pivotal.spring.cloud</groupId>
			<artifactId>spring-cloud-services-spring-connector</artifactId>
			<version>1.6.1.RELEASE</version>
		</dependency>		
	</dependencies>

	<dependencyManagement>
		<dependencies>
			<dependency>
				<groupId>org.springframework.cloud</groupId>
				<artifactId>spring-cloud-dependencies</artifactId>
				<version>${this.spring_cloud_version}</version>
				<type>pom</type>
				<scope>import</scope>
			</dependency>
		</dependencies>
	</dependencyManagement>

	<build>
		<plugins>
			<plugin>
				<groupId>org.springframework.boot</groupId>
				<artifactId>spring-boot-maven-plugin</artifactId>
			</plugin>
		</plugins>
	</build>


</project>
`;

    this.indexHtml = `
<!doctype html>
<html lang="en" ng-app="Microservice">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Microservice</title>

    <!-- Bootstrap core CSS -->
    <link href="webjars/bootstrap/4.0.0/css/bootstrap.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <style type="text/css">
        :root {
            --jumbotron-padding-y: 3rem;
        }

        .jumbotron {
            padding-top: var(--jumbotron-padding-y);
            padding-bottom: var(--jumbotron-padding-y);
            margin-bottom: 0;
            background-color: #fff;
        }
        @media (min-width: 768px) {
            .jumbotron {
                padding-top: calc(var(--jumbotron-padding-y) * 2);
                padding-bottom: calc(var(--jumbotron-padding-y) * 2);
            }
        }

        .jumbotron p:last-child {
            margin-bottom: 0;
        }

        .jumbotron-heading {
            font-weight: 300;
        }

        .jumbotron .container {
            max-width: 40rem;
        }
    </style>
</head>

<body>

<nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
    <a class="navbar-brand" href="#">Navbar</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
                <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#">Link</a>
            </li>
            <li class="nav-item">
                <a class="nav-link disabled" href="#">Disabled</a>
            </li>
            <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="http://example.com" id="dropdown01" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Dropdown</a>
                <div class="dropdown-menu" aria-labelledby="dropdown01">
                    <a class="dropdown-item" href="#">Action</a>
                    <a class="dropdown-item" href="#">Another action</a>
                    <a class="dropdown-item" href="#">Something else here</a>
                </div>
            </li>
        </ul>
        <form class="form-inline my-2 my-lg-0">
            <input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search">
            <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
        </form>
    </div>
</nav>

<main role="main">

    <section class="jumbotron text-center">
        <div class="container">
            <h1>Frontend UI sample template</h1>
            <p class="lead">microservice test sample.</p>
        </div>
    </section>

    <div class="album py-5 bg-light">
        <div class="container">
            <ng-view></ng-view>
        </div>
    </div>

</main>

<script type="text/javascript" src="webjars/jquery/2.2.4/jquery.min.js"></script>
<script type="text/javascript" src="webjars/bootstrap/4.0.0/js/bootstrap.js"></script>

<script type="text/javascript" src="webjars/angularjs/1.2.16/angular.js"></script>
<script type="text/javascript" src="webjars/angularjs/1.2.16/angular-resource.js"></script>
<script type="text/javascript" src="webjars/angularjs/1.2.16/angular-route.js"></script>
<script type="text/javascript" src="webjars/angular-ui/0.4.0/angular-ui.js"></script>

<script type="text/javascript" src="js/app.js"></script>
<script type="text/javascript" src="js/albums.js"></script>
<script type="text/javascript" src="js/errors.js"></script>

</body>
</html>
    
    `;
    this.albumsJs = `
angular.module('albums', ['ngResource']).
factory('Albums', function ($resource) {
    return $resource('api/search');
}).
factory('Albums2', function ($resource) {
    return $resource('api/searchTwo');
}).
factory('Albums3', function ($resource) {
    return $resource('api/searchThree');
}).
factory('Album', function ($resource) {
    return $resource('api/search/:id', {id: '@id'});
});

function AlbumsController($scope, Albums, Albums2, Albums3, Album) {
    function list() {
        $scope.albums = Albums.query();
        $scope.albums2 = Albums2.query();
        $scope.albums3 = Albums3.query();
        console.log($scope.albums)
    }

    function clone (obj) {
        return JSON.parse(JSON.stringify(obj));
    }

    $scope.init = function() {
        list();
    };
}



    
    `;
    this.appJs = `
angular.module('Microservice', ['albums', 'errors', 'ngRoute']).
config(function ($locationProvider, $routeProvider) {

        $routeProvider.when('/errors', {
            controller: 'ErrorsController',
            templateUrl: 'templates/errors.html'
        });
        $routeProvider.otherwise({
            controller: 'AlbumsController',
            templateUrl: 'templates/albums.html'
        });
    }
);    
    `;
    this.errorsJs = `
angular.module('errors', ['ngResource']).
factory('Errors', function ($resource) {
    return $resource('errors', {}, {
        kill: { url: 'errors/kill' },
        throw: { url: 'errors/throw' }
    });
});

function ErrorsController($scope, Errors, Status) {
    $scope.kill = function() {
        Errors.kill({},
            function () {
                Status.error("The application should have been killed, but returned successfully instead.");
            },
            function (result) {
                if (result.status === 502)
                    Status.error("An error occurred as expected, the application backend was killed: " + result.status);
                else
                    Status.error("An unexpected error occurred: " + result.status);
            }
        );
    };

    $scope.throwException = function() {
        Errors.throw({},
            function () {
                Status.error("An exception should have been thrown, but was not.");
            },
            function (result) {
                if (result.status === 500)
                    Status.error("An error occurred as expected: " + result.status);
                else
                    Status.error("An unexpected error occurred: " + result.status);
            }
        );
    };
}    
    `;
    this.albumsHtml = `
<div class="row" ng-init="init()">
    <div class="col-md-4">
        <div class="card mb-4 box-shadow">
            <div class="card-body">
                <p class="card-text">Backend Service API Call - 1</p>
                <div class="d-flex justify-content-between align-items-center">
                    <div class="btn-group">
                        <button type="button" class="btn btn-sm btn-outline-secondary" data-toggle="collapse" href="#collapseExample">View</button>
                    </div>
                </div>
                <p>
                <div class="collapse" id="collapseExample">
                    <div class="card card-body">
                        {{albums}}
                    </div>
                </div>
                </p>
            </div>
        </div>
    </div>
    <div class="col-md-4">
        <div class="card mb-4 box-shadow">
            <div class="card-body">
                <p class="card-text">BackEnd Service API Call - 2</p>
                <div class="d-flex justify-content-between align-items-center">
                    <div class="btn-group">
                        <button type="button" class="btn btn-sm btn-outline-secondary" data-toggle="collapse" href="#collapseExample2">View</button>
                    </div>
                </div>
                <p>
                <div class="collapse" id="collapseExample2">
                    <div class="card card-body">
                        {{albums2}}
                    </div>
                </div>
                </p>
            </div>
        </div>
    </div>
    <div class="col-md-4">
        <div class="card mb-4 box-shadow">
            <div class="card-body">
                <p class="card-text">BackEnd Service API Call - 3</p>
                <div class="d-flex justify-content-between align-items-center">
                    <div class="btn-group">
                        <button type="button" class="btn btn-sm btn-outline-secondary"  data-toggle="collapse" href="#collapseExample3">View</button>
                    </div>
                </div>
                <p>
                <div class="collapse" id="collapseExample3">
                    <div class="card card-body">
                        {{albums3}}
                    </div>
                </div>
                </p>
            </div>
        </div>
    </div>


</div>   
    `;
    this.errorsHtml = `
<div id="errors" class="col-xs-12"  ng-controller="ErrorsController">
    <div class="page-header">
        <h1>Force Errors</h1>
    </div>

    <div class="row">
        <form role="form">
            <div class="form-group">
                <h3>Kill this instance of the application</h3>
                <a ng-click="kill()" class="btn btn-primary btn-lg active" role="button">Kill</a>
            </div>
            <div class="form-group">
                <h3>Force an exception to be thrown from the application</h3>
                <a ng-click="throwException()" class="btn btn-primary btn-lg active" role="button">Throw Exception</a>
            </div>
        </form>
    </div>
</div>    
    `;

    this.frontManifest = `---
applications:
  - name: front
    memory: 1G
    path: target/front-0.0.1-SNAPSHOT.jar
    env:
      msa: yes   
    `;

    this.backManifest = `---
applications:
  - name: back
    memory: 1G
    path: target/back-0.0.1-SNAPSHOT.jar
    env:
      msa: yes    
    `;
  }
}
