import { IMediaResource } from '../media'
import { IMemberResource } from './member'
import { IBranchResource } from '../branch'
import { IEmployeeResource } from '../employee'
import { IMemberTypeResource } from './member-type'
import { IMemberAssetsResource } from './member-assets'
import { IMemberWalletResource } from './member-wallet'
import { IMemberIncomeResource } from './member-income'
import { IMemberGenderResource } from './member-gender'
import { IMemberCenterResource } from './member-center'
import { IMemberAddressResource } from './member-address'
import { IMemberExpensesResource } from './member-expenses'
import { IMemberRecruitsResource } from './member-recruits'
import { IMemberOccupationResource } from './member-occupation'
import { IMemberDescriptionResource } from './member-description'
import { IMemberCloseRemarksResource } from './member-close-remarks'
import { IMemberJointAccountsResource } from './member-joint-accounts'
import { IMemberClassificationResource } from './member-classification'
import { IMemberRelativeAccountsResource } from './member-relative-accounts'
import { IMemberGovernmentBenefitsResource } from './member-government-benefits'
import { ITimeStamps, TAccountStatus, TCivilStatus, TEntityId } from '../common'
import { IMemberMutualFundsHistoryResource } from './member-mutual-funds-history'
import { IMemberContactNumberReferencesResource } from './member-contact-number-references'

export interface IMemberProfileRequest {
    oldReferenceId?: string
    passbookNumber?: string

    notes: string
    description: string
    contactNumber: string
    tinNumber?: string
    civilStatus: TCivilStatus
    occupationId: TEntityId
    sssNumber?: string
    businessAddress?: string
    businessContact?: string

    status: 'Pending' | 'Verified' | 'Not Allowed'
    isClosed: boolean

    pagibigNumber?: string
    philhealthNumber?: string
    isMutualFundMember: boolean
    isMicroFinanceMember: boolean

    mediaId?: TEntityId
    memberId?: TEntityId

    branchId?: TEntityId
    memberTypeId?: TEntityId
    memberClassificationId?: TEntityId
    memberGenderId?: TEntityId
    memberCenterId?: TEntityId
    memberEducationalAttainmentId?: TEntityId

    memberDescriptions?: { name: string; description: string }[]

    memberAddress: {
        postalCode: string
        province: string
        city: string
        barangay: string
        label: string
    }[]

    memberContactNumberReferences: {
        name: string
        description: string
        contactNumber: string
    }[]

    memberIncome?: {
        name: string
        amount: number
        date: string
        description: string
    }[]

    memberRelativeAccounts?: {
        membersProfileId: TEntityId
        relativeProfileMemberId: TEntityId
        familyRelationship: string
        description: string
    }[]

    memberAssets?: {
        entryDate: string
        description: string
        name: string
    }[]

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
        membersProfileId: TEntityId
        membersProfileRecruitedId: TEntityId
        dateRecruited: string
        description: string
        name: string
    }[]

    memberGovernmentBenefits?: {
        country: string
        name: string
        description: string
        value: string
        frontMediaId?: TEntityId
        backMediaId?: TEntityId
    }[]
}

export interface IMemberProfileResource extends ITimeStamps {
    id: TEntityId
    oldReferenceId?: string
    passbookNumber?: string

    notes: string
    description: string
    contactNumber: string
    tinNumber?: string
    civilStatus: string
    sssNumber?: string
    businessAddress?: string
    businessContact?: string

    status: TAccountStatus
    isClosed: boolean

    pagibigNumber?: string
    philhealthNumber?: string
    isMutualFundMember: boolean
    isMicroFinanceMember: boolean

    occupationId?: TEntityId
    occupation: IMemberOccupationResource

    mediaId?: TEntityId
    media?: IMediaResource

    memberId?: TEntityId
    member?: IMemberResource

    memberTypeId?: TEntityId
    memberType?: IMemberTypeResource

    memberClassificationId?: TEntityId
    memberClassification?: IMemberClassificationResource

    memberGenderId?: TEntityId
    memberGender?: IMemberGenderResource

    verifiedByEmployeeId?: TEntityId
    verifiedByEmployee?: IEmployeeResource

    branchId?: TEntityId
    branch?: IBranchResource

    memberCenterId?: TEntityId
    memberCenter?: IMemberCenterResource

    signatureMediaId?: TEntityId
    signatureMedia?: IMediaResource

    memberEducationalAttainmentId?: TEntityId
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
