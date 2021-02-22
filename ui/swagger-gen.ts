// tslint:disable
/* eslint-disable */
const { codegen } = require('swagger-axios-codegen');
const swaggerLocation = "../docs/swagger.json"

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/views/route/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['Route'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/views/system/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['System'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/views/namespace/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['Namespace'],
});

codegen({
  methodNameMode: 'operationId',
  source: require(swaggerLocation),
  outputDir: './src/views/tag/generated/',
  useCustomerRequestInstance: true,
  serviceNameSuffix: 'Service',
  useStaticMethod: false,
  include: ['Tag'],
});

