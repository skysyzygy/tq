# tq - json 🔀 tessitura 

![go workflow](https://github.com/skysyzygy/tq/actions/workflows/go.yml/badge.svg)
[![codecov](https://codecov.io/gh/skysyzygy/tq/graph/badge.svg?token=Ov8gpBWhHQ)](https://codecov.io/gh/skysyzygy/tq)

**tq** is a wrapper around the Tessitura API that reads JSON-formatted data and outputs a series of API calls to Tessitura. It internally handles authentication, session creation and closure, and batch/concurrent processing so that humans like you (or bots or scripts) can focus on the data and not the intricacies of the Tessitura API.                                                       
                     
**tq** is basically a high-level API for common tasks in Tessi.

## usage: 

``` tq [flags] [verb] [object] [query]```

### flags:
* **-n, --dryrun**         : don't actually do anything, just show what would have happened
* **-f, --file** *string*  : JSON file to read (default is to read from stdin)
* **-h, --help**           : help for tq
* **-l, --log** *string*   : log file to write to (default is no log)
* **-v, --verbose**        : turns on additional diagnostic output

#### configuration file:
A yaml configuration file `.tq` placed in your home directory can be used to set defaults for these flags; it is also used to save the current authentication method. See `tq auth select --help` for more information. 


### verbs:
*  **authenticate** : Authenticate with the Tessitura API
*  **create** :       Create entities in Tessitura
*  **get** :          Retrieve entities from Tessitura
*  **update** :       Update entities in Tessitura
*  **completion** :   Generate the autocompletion script for the specified shell
*  **help** :         Help about any command

### queries:
Queries are simply JSON objects and can be batched by combining multiple query objects into a single JSON array, e.g. 

```[{"ID":123},{"ID":124},{"ID":125},...]```

Query details are detailed in the help for each command.


### objects:
*  **AccountTypes** :                            Get the details of an account type by id
*  **Accounts** :                                Get details of a specific credit card account
*  **ActionTypes** :                             Get the details of an action type by id
*  **Actions** :                                 Get details of an issue action
*  **ActivityCategories** :                      Get the details of an activity category by id
*  **ActivityTypes** :                           Get the details of an activity type by id
*  **AddressTypes** :                            Get the details of an address type by id
*  **Addresses** :                               Get details of an address using addressId as a URL query parameter
*  **AffiliationTypes** :                        Get the details of an affiliation type by id
*  **Affiliations** :                            Get details of an affiliation
*  **AliasTypes** :                              Get the details of an alias type by id
*  **Aliases** :                                 Get details of an alias
*  **AnalyticsCubes** :                          Get the details of an analytics cube
*  **AnalyticsReports** :                        Get a single SSRS Report for display in Analytics
*  **AppScreenTexts** :                          Get the details of an App Screen Text by id
*  **AppealCategories** :                        Get the details of an appeal category by id
*  **Appeals** :                                 Get details of an Appeal
*  **ApplicationObjects** :                      Get all application objects valid for the context usergroup
*  **Artists** :                                 Get details of an existing artist
*  **AssetTypes** :                              Get the details of an asset type by id
*  **Assets** :                                  Get details of an asset
*  **AssociationTypes** :                        Get the details of an association type by id
*  **Associations** :                            Get details of an association
*  **AttendanceHistory** :                       Attendance History for a selected constituent optionally including primary affiliates
*  **Attributes** :                              Get details of an attribute
*  **AuditLogs** :                               Get details of a audit log
*  **Authenticate** :                            This is a no-op operation for windows authentication diagnostics
*  **BatchMaintenance** :                        Get a single Batch
*  **BatchTypeGroups** :                         Get the details of a batch type group by id
*  **BatchTypeUserGroup** :                      Get all batch type/user group mappings
*  **BatchTypes** :                              Get the details of a batch type by id
*  **BillingSchedules** :                        Get the details of a Billing Schedule
*  **BillingTypes** :                            Get the details of a Billing Type by id
*  **BookingCategories** :                       Get the details of a Booking Category by id
*  **BookingTemplates** :                        Get a Booking Template by ID
*  **Bookings** :                                Get a Booking by id
*  **BulkCopySets** :                            Get a bulk copy set by Id
*  **BulkDailyCopyExclusions** :                 Get a bulk daily copy exclusion by id
*  **BusinessUnits** :                           Get the details of a business unit by id
*  **Cache** :                                   
*  **CampaignDesignations** :                    Get a single Designation associated to a Campaign
*  **CampaignFunds** :                           Get a single Fund associated to a Campaign
*  **Campaigns** :                               Get summary of a specific campaign
*  **CardReaderTypes** :                         Get the details of a Card Reader Type by id
*  **Cart** :                                    Gets the cart details
*  **Colors** :                                  Get the details of a color by id
*  **Composers** :                               Get the details of a composer by id
*  **Constituencies** :                          Get details of constituency
*  **ConstituencyTypes** :                       Get the details of a constituency type by id
*  **ConstituentContributions** :                Get contributions for a constituent
*  **ConstituentDocuments** :                    Get the details of a document for a constituent
*  **ConstituentGroups** :                       Get the details of a constituent group by id
*  **ConstituentInactives** :                    Get the details of a constituent inactive by id
*  **ConstituentProtectionTypes** :              Get the details of a constituent protection type by id
*  **ConstituentTypeAffiliates** :               Get the details of a constituent type affiliate by id
*  **ConstituentTypes** :                        Get the details of a constituent type by id
*  **Constituents** :                            Get the details of a Constituent using id
*  **ContactPermissionCategories** :             Get the details of a contact permission category
*  **ContactPermissionTypes** :                  Get the details of a contact permission type
*  **ContactPermissions** :                      Get details of a contact permission
*  **ContactPointCategories** :                  Get the details of a contact point category by id
*  **ContactPointCategoryPurposes** :            Get the details of a contact point category purpose by id
*  **ContactPointPurposeCategories** :           Get the details of a contact point purpose category by id
*  **ContactPointPurposeMaps** :                 Get details of a contact point purpose
*  **ContactPointPurposes** :                    Get the details of a contact point purpose by id
*  **ContactPoints** :                           Get all the delivery points for the specified constituent (constituentId) and all its visible affiliation's delivery point purposes as well
*  **ContactTypes** :                            Get the details of a contact type by id
*  **ContextInformation** :                      Get a commonly used set of default values for the user and usergroup in the current security context
*  **ContributionDesignations** :                Get the details of a contribution designation by id
*  **ContributionImportSets** :                  Get the details of a contributionImportSet by id
*  **ControlGroupUserGroups** :                  Get the details of a control group/user group mapping by id
*  **ControlGroups** :                           Get the details of a control group by id
*  **CoreIdentity** :                            
*  **Countries** :                               Get the details of a country by id
*  **CrediteeTypes** :                           Get the details of a crediteeType by id
*  **Credits** :                                 Returns all credits for the requested production element
*  **CriterionOperators** :                      Get the details of a criterion operator by id
*  **CumulativeGivingReceipts** :                Get details of an cumulativeGivingReceipt
*  **CurrencyTypes** :                           Get the details of a currency type by id
*  **Custom** :                                  Get details of an entry in the table for the resource as defined by {resourceName} in TR_DATASERVICE_TABLES with the given id {Id}
*  **CustomDefaultCategories** :                 Get the details of a custom default category by id
*  **CustomDefaults** :                          Get the details of a custom default by id
*  **DeliveryMethods** :                         Get the details of a delivery method by id
*  **DesignationCodes** :                        Get the details of a designation code by id
*  **Designs** :                                 Get the details of a design by id
*  **Diagnostics** :                             Validates Encryption Key Dates
*  **DirectDebitAccountTypes** :                 Get the details of a direct debit account type by id
*  **DiscountTypes** :                           Get the details of a discount type by id
*  **Divisions** :                               Get the control group/division mappings for the specified division
*  **DocumentCategories** :                      Get the details of a documentCategory by id
*  **Documents** :                               Get the details of a document
*  **DonationLevels** :                          Get the details of a donation level by id
*  **EMV** :                                     Retrieve information on all lanes associated with merchant
*  **ElectronicAddressTypes** :                  Get the details of an electronic address type by id
*  **ElectronicAddresses** :                     Get details of an electronic address
*  **EmailProfiles** :                           Get the details of an email profile by id
*  **EmarketIndicators** :                       Get the details of an emarket indicator by id
*  **Eras** :                                    Get the details of an era by id
*  **EventControl** :                            Returns a response containing a list of EventControl rows for the N-Scan event control table
*  **Facilities** :                              Get details of a Facility
*  **Fees** :                                    Get details of a fee
*  **FinanceContributions** :                    Get details of a contribution
*  **Formats** :                                 Get the details of a format by id
*  **Funds** :                                   Get details of a specific fund
*  **GLAccounts** :                              Get the details of a gl account by id
*  **Genders** :                                 Get the details of a gender by id
*  **GiftAidContactMethods** :                   Get the details of a gift aid contact method by id
*  **GiftAidDeclarations** :                     Gets a single Gift Aid Declaration
*  **GiftAidDocumentStatuses** :                 Get the details of a gift aid document status by id
*  **GiftAidIneligibleReasons** :                Get the details of a gift aid ineligible reason by id
*  **GiftAidRates** :                            Get the details of a gift aid rate by id
*  **GiftAidStatuses** :                         Get the details of a gift aid status by id
*  **GiftAidTypes** :                            Get the details of a gift aid type by id
*  **GiftCertificates** :                        Get details of a specific gift certificate
*  **HoldCodeCategories** :                      Get the details of a Hold Code Category by id
*  **HoldCodeUserGroups** :                      Get the details of a hold code/user group mapping by id
*  **HoldCodes** :                               Get details of a Hold Code
*  **InactiveReasons** :                         Get the details of an inactive reason by id
*  **IntegrationDefaults** :                     Get the details of an Integration Default by id
*  **Integrations** :                            Get the details of an action type by id
*  **InterestCategories** :                      Get the details of an interest category by id
*  **InterestTypes** :                           Get the details of an interest type by id
*  **Interests** :                               Get details of an interest
*  **Internal** :                                Get details of an address and all the phones attached to it
*  **InventoryContactPermissionTypes** :         Get details of an inventoryContactPermissionType
*  **InventoryWebContents** :                    Get details of an inventoryWebContent
*  **InvoiceBilling** :                          Get status of a processing automated billing run
*  **Issues** :                                  Get details of a customer issue
*  **KeywordCategories** :                       Get the details of a keyword category by id
*  **Keywords** :                                Get the details of a keyword by id
*  **Languages** :                               Get the details of a language by id
*  **ListCategories** :                          Get the details of a list category by id
*  **Lists** :                                   Get a single List without contents
*  **LoginTypes** :                              Get the details of a login type by id
*  **MachineSettings** :                         Get the details of a Machine Setting by id
*  **MailIndicators** :                          Get the details of a mail indicator by id
*  **MediaTypes** :                              Get the details of a media type by id
*  **MembershipLevelCategories** :               Get the details of a membership level category by id
*  **MembershipLevels** :                        Get a specific membership level
*  **MembershipOrganizations** :                 Get a specific membership organization
*  **MembershipStandings** :                     Get a specific MembershipStanding by id
*  **Memberships** :                             Returns membership details for a constituent
*  **MerchantReferences** :                      
*  **Merchants** :                               Retrieve Merchant List
*  **ModeOfSaleCategories** :                    Get details of a mode of sale category
*  **ModeOfSaleOffers** :                        Get details of a mode of sale offer
*  **ModeOfSalePriceTypes** :                    Get details of a mode of sale price type
*  **ModeOfSaleSurveyQuestions** :               Get details of a mode of sale survey question
*  **ModeOfSaleUserGroups** :                    Get the details of a mode of sale/user group mapping by id
*  **ModesOfSale** :                             Get details of an existing mode of sale
*  **NScanAccessAreas** :                        Get the details of an NScan Access Area by id
*  **NameStatuses** :                            Get the details of a name status by id
*  **ObjectPermissions** :                       Get the details of an object permission by id
*  **OrderBilling** :                            Reprint Order billing
*  **OrderCategories** :                         Get the details of an order category by id
*  **Orders** :                                  Get the details of an existing Order as the contract OrderProductsView
*  **Organizations** :                           Get the details of an organization by ID
*  **OriginalSources** :                         Get the details of an original source by id
*  **Origins** :                                 Get the details of an origin by id
*  **OutputSets** :                              Get a single Output Set
*  **PackageHistory** :                          Gets history information for packages
*  **PackagePriceTypes** :                       Get details of a specific package price type
*  **PackageTypes** :                            Get the details of a package type by id
*  **PackageWebContents** :                      Get details of a packageWebContent
*  **Packages** :                                Get details of a package
*  **PaymentGatewayActivities** :                Get a single Payment Gateway Activity
*  **PaymentGatewayConfiguration** :             Retrieve Payment Gateway Configuration
*  **PaymentGatewayNotifications** :             Get all notification events by reference
*  **PaymentGatewayTransactionTypes** :          Get the details of a Payment Gateway Transaction Type by id
*  **PaymentHistory** :                          Gets history information for payments by constituent
*  **PaymentMethodGroups** :                     Get the details of a payment method group by id
*  **PaymentMethodUserGroups** :                 Get details of a payment method/user group mapping by id
*  **PaymentMethods** :                          Get details of a payment method
*  **PaymentSignatures** :                       Get details of a payment signature
*  **PaymentTypes** :                            Get the details of a payment type by id
*  **Payments** :                                Get a single payment
*  **PerformanceGroups** :                       Get details of a Performance Group
*  **PerformancePackageModeOfSales** :           Get details of a performance package mode of sale
*  **PerformancePriceLayers** :                  Get details of a performance price layer
*  **PerformancePriceTypes** :                   Get details of a performance price type
*  **PerformancePrices** :                       Get details of a performance price
*  **PerformanceStatuses** :                     Get the details of a performance status by id
*  **PerformanceTypes** :                        Get the details of a performance type by id
*  **Performances** :                            Get details of a performance
*  **Philanthropy** :                            Get details of an philanthropyEntry
*  **PhilanthropyTypes** :                       Get the details of philosophy type by id
*  **PhoneIndicators** :                         Get the details of a phone indicator by id
*  **PhoneTypes** :                              Get the details of a phone type by id
*  **Phones** :                                  Get details of a phone
*  **PlanPriorities** :                          Get the details of a plan priority by id
*  **PlanSources** :                             Get the details of a plan source by id
*  **PlanStatuses** :                            Get the details of a plan status by id
*  **PlanTypes** :                               Get the details of a plan type by id
*  **PlanWorkers** :                             Get details of a plan worker
*  **Plans** :                                   Get details of a plan
*  **PledgeBilling** :                           Get status of a pledge billing run
*  **PortfolioCustomElements** :                 Get the details of a portfolio custom element by id
*  **Portfolios** :                              Get portfolio for a constituent
*  **Prefixes** :                                Get the details of a prefix by id
*  **Premieres** :                               Get the details of a premiere by id
*  **PriceCategories** :                         Get the details of a price category by id
*  **PriceEvents** :                             Get details of a price event
*  **PriceLayerTypes** :                         Get the details of a price layer type by id
*  **PriceTemplates** :                          Get details of a price template
*  **PriceTypeCategories** :                     Get the details of a price type category by id
*  **PriceTypeGroups** :                         Get the details of a price type group by id
*  **PriceTypeReasons** :                        Get the details of a price type reason by id
*  **PriceTypeUserGroups** :                     Get details of a price type/user group mapping by id
*  **PriceTypes** :                              Get details of a price type
*  **PricingRuleCategories** :                   Get the details of a pricing rule category by id
*  **PricingRuleMessageTypes** :                 Get the details of a pricing rule message type by id
*  **PricingRuleSets** :                         Get details of a pricing rule set
*  **PricingRuleTypes** :                        Get the details of a pricing rule type by id
*  **PricingRules** :                            Get details of a pricing rule
*  **Printers** :                                Get the details of a printer by id
*  **ProductKeywords** :                         Returns keywords for the requested production elements or packages
*  **ProductionSeasonMembershipOrganizations** : Get details of a specific production season membership organization
*  **ProductionSeasons** :                       Get details of a specific production season
*  **Productions** :                             Get details of a specific production
*  **ProgramListings** :                         Get details of a program listing
*  **Programs** :                                Get the details of a program by id
*  **Pronouns** :                                
*  **QualificationCategories** :                 Get the details of a Qualification Category by id
*  **Qualifications** :                          Get the details of a Qualification by id
*  **QueryElementFilters** :                     Get a specific query element filter
*  **QueryElementGroups** :                      Get the details of a query element group by id
*  **QueryElements** :                           Get details of a query element by id
*  **RankTypes** :                               Get the details of a rank type by id
*  **Rankings** :                                Get details of a ranking
*  **ReceiptSettings** :                         Get the details of a Receipt Setting by id
*  **ReferenceColumns** :                        Get all reference columns
*  **ReferenceTableUserGroups** :                Get the details of a reference table/user group mapping by id
*  **ReferenceTables** :                         Get details for a reference table by Id
*  **RelationshipCategories** :                  Get the details of a relationship category by id
*  **Relationships** :                           Get all affiliations and associations of the specified constituent id or get all affiliations and associations of the specified associated constituent id
*  **ReportRequests** :                          Get a report request
*  **ReportSchedules** :                         Get a report schedule
*  **ReportUserGroups** :                        Get the details of a report/user group mapping by id
*  **Reports** :                                 Get details of a report
*  **Research** :                                Get details of a research entry
*  **ResearchTypes** :                           Get the details of research type by id
*  **ResourceCategories** :                      Get the details of a Resource Category by id
*  **ResourceSchedules** :                       Get a single resource schedule
*  **ResourceTypes** :                           Get all resource types
*  **Resources** :                               Get a specific Resource
*  **SalesChannels** :                           Get the details of a sales channel by id
*  **SalesLayoutButtonTypes** :                  Get the details of a sales layout button type by id
*  **SalesLayouts** :                            Get details of a sales layout setup
*  **SalutationTypes** :                         Get the details of a salutation type by id
*  **Salutations** :                             Get details of a salutation
*  **SchedulePatternTypes** :                    Get the details of a Schedule Pattern by id
*  **ScheduleTypes** :                           Get the details of a Schedule Type by id
*  **SeasonTypes** :                             Get the details of a season type by id
*  **Seasons** :                                 Get the details of a season by id
*  **SeatCodes** :                               Get the details of a seat code by id
*  **SeatStatuses** :                            Get the details of a seat status by id
*  **Sections** :                                Get the details of a section by id
*  **SecurityBatchTypes** :                      Get all batch type/user group mappings valid for the context usergroup
*  **SecurityControlGroups** :                   Get all control group/user group mappings valid for the context usergroup
*  **SecurityHoldCodes** :                       Get all hold code/user group mappings valid for the context usergroup
*  **SecurityModesOfSale** :                     Get all mode of sale/user group mappings valid for the context usergroup
*  **SecurityObjectPermissions** :               Get all the object permissions valid for the context usergroup
*  **SecurityPaymentMethods** :                  Get all payment method/user group mappings valid for the context usergroup
*  **SecurityPriceTypes** :                      Get all price type/user group mappings valid for the context usergroup
*  **SecurityReferenceTables** :                 Get all the reference table/user group mappings valid for the context usergroup
*  **SecurityReports** :                         Get all report/user group mappings valid for the context usergroup
*  **SecurityServiceResources** :                Get all service resource/user group mappings valid for the context usergroup
*  **SecurityUserGroups** :                      Get all security user groups
*  **ServiceResourceUserGroups** :               Get the details of a service resource/user group mapping by id
*  **ServiceResources** :                        Get all service resources
*  **Session** :                                 Returns details summarizing a web session's state
*  **SourceGroups** :                            Get the details of a source group by id
*  **Sources** :                                 Get details of a Source
*  **SpecialActivities** :                       Get details of an activity record
*  **SpecialActivityStatuses** :                 Get the details of a Special Activity Status by id
*  **SpecialActivityTypes** :                    Get the details of a Special Activity Type by id
*  **States** :                                  Get the details of a state by id
*  **StepTypes** :                               Get the details of a step type by id
*  **Steps** :                                   Get details of a step
*  **SubLineItemStatuses** :                     Get the details of a sub line item status by id
*  **SubLineItems** :                            Returns sub line item summary data for a constituent
*  **Suffixes** :                                Get the details of a suffix by id
*  **SurveyQuestions** :                         Get the details of a survey question by id
*  **SurveyResponses** :                         Get details of a survey response
*  **SystemDefaults** :                          Get all system defaults
*  **TemplateCategories** :                      Get the details of a template category by id
*  **TemplatePriceTypes** :                      Get details of a template price type
*  **TemplatePrices** :                          Get details of a template price
*  **TemplateTypes** :                           Get the details of a template type by id
*  **Templates** :                               Get the details of a template by id
*  **Theaters** :                                Get the details of a theater by id
*  **TicketHistory** :                           Gets history information for tickets
*  **TimeSlots** :                               Get the details of a time slot by id
*  **Titles** :                                  Get details of a specific title
*  **TransactionHistory** :                      Get details of all transaction histories for the specified constituent id and all its visible affiliations' transaction histories as well
*  **TriPOSCloudConfigurations** :               Get the details of a TriPOS Cloud configuration by id
*  **UpgradeCategories** :                       Get the details of an Upgrade Category by id
*  **UpgradeLogs** :                             Returns an upgradeLog for the given id
*  **UserGroups** :                              Get the details of a user group by id
*  **UserPreferences** :                         Get a specific user preference by key
*  **Users** :                                   Get the details of a user for the specified username
*  **WebContentTypes** :                         Get a specific web content type
*  **WebContents** :                             Returns web content for the requested production elements or packages
*  **WebLogins** :                               Get details of a weblogin
*  **WorkerQualifications** :                    Get a single WorkerQualification by Id
*  **WorkerRoles** :                             Get the details of a worker role by id
*  **WorkerTypes** :                             Get the details of a worker type by id
*  **Workers** :                                 Get details of a worker
*  **ZoneGroups** :                              Get the details of a zone group by id
*  **ZoneMaps** :                                Get details of a specific zone map
*  **Zones** :                                   Get details of a specific zone