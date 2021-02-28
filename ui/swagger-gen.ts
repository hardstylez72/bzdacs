// tslint:disable
/* eslint-disable */
const { codegen } = require('swagger-axios-codegen');
const swaggerLocation = "../generated/swagger.json"

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/route/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['Route'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/system/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['System'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/namespace/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['Namespace'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/tag/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['Tag'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/group/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['Group', 'GroupRoute'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/user/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['User', 'UserGroup', 'UserRoute'],
});



