export class Tree {
  dirs: object;

  constructor() {
    this.dirs = { 'core' : {
      'data' : [
        {
          'text' : 'frontend web',
          'children' : [
            {
              'text' : 'src',
              'children' : [
                {
                  'text':'main',
                  'children' : [
                    {
                      'text':'java',
                      'children':[
                        {
                          'text':'com',
                          'children':[
                            {
                              'text':'crossent',
                              'children':[
                                {
                                  'text':'microservice',
                                  'children':[
                                    {'id':'front_data','text':'Data.java'}, {'text':'FrontApplication.java'}, {'text':'FrontController.java'}
                                  ]
                                },
                              ]
                            }
                          ]
                        }

                      ]
                    },
                    {
                      'text':'resources',
                      'children':[
                        {'id':'front_application_properties','text':'application.properties'},{'id':'front_bootstrap_properties','text':'bootstrap.properties'},
                        {'text':'static',
                          'children':[
                            {'id':'index.html','text':'index.html'},
                            {'text':'js',
                              'children':[
                                {'id':'albums.js','text':'albums.js'},{'id':'app.js','text':'app.js'},{'id':'errors.js','text':'errors.js'}
                              ]
                            },
                            {'text':'templates',
                              'children':[
                                {'id':'albums.html','text':'albums.html'}, {'id':'errors.html','text':'errors.html'}
                              ]
                            }

                          ]
                        }
                      ]
                    }
                  ]
                }
              ]
            },
            {'id':'front_pom','text':'pom.xml'},{'id':'front_manifest','text':'manifest.yml'}
          ]
        },
        {
          'text' : 'backend service',
          'children' : [
            {
              'text' : 'src',
              'children' : [
                {
                  'text':'main',
                  'children' : [
                    {
                      'text':'java',
                      'children':[
                        {
                          'text':'com',
                          'children':[
                            {
                              'text':'crossent',
                              'children':[
                                {
                                  'text':'microservice',
                                  'children':[
                                    {'id':'back_data','text':'Data.java'}, {'text':'BackApplication.java'}, {'text':'BackController.java'}
                                  ]
                                },
                              ]
                            }
                          ]
                        }

                      ]
                    },
                    {
                      'text':'resources',
                      'children':[
                        {'id':'back_application_properties','text':'application.properties'},{'id':'back_bootstrap_properties','text':'bootstrap.properties'}
                      ]
                    }
                  ]
                }
              ]
            },
            {'id':'back_pom','text':'pom.xml'},{'id':'back_manifest','text':'manifest.yml'}
          ]
        }
      ]
    } };
  }

}
