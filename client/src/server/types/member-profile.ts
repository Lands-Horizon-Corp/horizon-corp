import { IBranchResource } from './branch'
import { IMemberResource } from './member'
import { IEmployeeResource } from './employee'
import { ITimeStamps, TAccountStatus, TEntityId } from './common'
import { IMemberTypeResource } from './member-type'
import { IMemberAssetsResource } from './member-assets'
import { IMemberWalletResource } from './member-wallet'
import { IMemberIncomeRequest, IMemberIncomeResource } from './member-income'
import { IMemberGenderResource } from './member-gender'
import { IMemberCenterResource } from './member-center'
import { IMemberAddressRequest, IMemberAddressResource } from './member-address'
import { IMemberExpensesResource } from './member-expenses'
import { IMemberRecruitsResource } from './member-recruits'
import {
    IMemberDescriptionRequest,
    IMemberDescriptionResource,
} from './member-description'
import {
    IMemberDescriptionRequest,
    IMemberDescriptionResource,
} from './member-description'
import { IMemberCloseRemarksResource } from './member-close-remarks'
import { IMemberJointAccountsResource } from './member-joint-accounts'
import { IMemberClassificationResource } from './member-classification'
import { IMemberRelativeAccountsResource } from './member-relative-accounts'
import { IMemberGovernmentBenefitsResource } from './member-government-benefits'
import { IMemberMutualFundsHistoryResource } from './member-mutual-funds-history'
import {
    IMemberContactNumberReferencesRequest,
    IMemberContactNumberReferencesResource,
} from './member-contact-number-references'
import { IMediaResource } from './media'

export interface IMemberProfileRequest {
    oldReferenceID?: string
    passbookNumber?: string

    notes: string
    description: string
    contactNumber: string
    tinNumber?: string
    civilStatus: string
    occupation?: string
    sssNumber?: string
    businessAddress?: string
    businessContact?: string

    status: TAccountStatus
    isClosed: boolean

    pagibigNumber?: string
    philhealthNumber?: string
    isMutualFundMember: boolean
    isMicroFinanceMember: boolean

    mediaID?: TEntityId
    memberID?: TEntityId

    memberTypeID?: TEntityId
    memberClassificationID?: TEntityId
    memberGenderID?: TEntityId
    branchID?: TEntityId
    memberCenterID?: TEntityId
    memberEducationalAttainmentID?: TEntityId

    memberAddresses?: IMemberAddressRequest[]
    memberDescriptions?: IMemberDescriptionRequest[]
    memberContactNumberReferences?: IMemberContactNumberReferencesRequest[]
    memberIncome?: IMemberIncomeRequest[]

    memberRelativeAccounts?: {
        membersProfileID: TEntityId
        relativeProfileMemberID: TEntityId
        familyRelationship: string
        description: string
    }[]
    memberAssets?: { entryDate: string; description: string; name: string }[]

    memberExpenses?: {
        name: string
        date: string
        amount: number
        description: string
    }[]
    memberJointAccounts?: {
        description: string
        firstName: string
        lastName: string
        middleName?: string
        familyRelationship?: string
    }[]
    memberRecruits?: {
        membersProfileID: TEntityId
        membersProfileRecruitedID: TEntityId
        dateRecruited: string
        description: string
        name: string
    }[]

    memberGovernmentBenefits?: {
        country: string
        name: string
        description: string
        value: number
        frontMediaID?: TEntityId
        backMediaID?: TEntityId
    }[]

    // memberWallets?: IMemberWalletResource[]
    // memberCloseRemarks?: IMemberCloseRemarksResource[]
    // memberMutualFundsHistory?: IMemberMutualFundsHistoryResource[]
}

export interface IMemberProfileResource extends ITimeStamps {
    id: TEntityId
    oldReferenceID?: string
    passbookNumber?: string

    notes: string
    description: string
    contactNumber: string
    tinNumber?: string
    civilStatus: string
    occupation?: string
    sssNumber?: string
    businessAddress?: string
    businessContact?: string

    status: TAccountStatus
    isClosed: boolean

    pagibigNumber?: string
    philhealthNumber?: string
    isMutualFundMember: boolean
    isMicroFinanceMember: boolean

    mediaID?: TEntityId
    media?: IMediaResource

    memberID?: TEntityId
    member?: IMemberResource

    memberTypeID?: TEntityId
    memberType?: IMemberTypeResource

    memberClassificationID?: TEntityId
    memberClassification?: IMemberClassificationResource

    memberGenderID?: TEntityId
    memberGender?: IMemberGenderResource

    verifiedByEmployeeID?: TEntityId
    verifiedByEmployee?: IEmployeeResource

    branchID?: TEntityId
    branch?: IBranchResource

    memberCenterID?: TEntityId
    memberCenter?: IMemberCenterResource

    signatureMediaID?: TEntityId
    signatureMedia?: IMediaResource

    memberEducationalAttainmentID?: TEntityId
    memberDescriptions?: IMemberDescriptionResource[]
    memberRecruits?: IMemberRecruitsResource[]
    memberContactNumberReferences?: IMemberContactNumberReferencesResource[]
    memberWallets?: IMemberWalletResource[]
    memberIncome?: IMemberIncomeResource[]
    memberExpenses?: IMemberExpensesResource[]
    memberCloseRemarks?: IMemberCloseRemarksResource[]
    memberJointAccounts?: IMemberJointAccountsResource[]
    memberRelativeAccounts?: IMemberRelativeAccountsResource[]
    memberAddresses?: IMemberAddressResource[]
    memberGovernmentBenefits?: IMemberGovernmentBenefitsResource[]
    memberMutualFundsHistory?: IMemberMutualFundsHistoryResource[]
    memberAssets?: IMemberAssetsResource[]
}
