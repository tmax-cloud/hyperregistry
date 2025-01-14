import {ComponentFixture, ComponentFixtureAutoDetect, TestBed, waitForAsync} from '@angular/core/testing';
import {TranslateService} from '@ngx-translate/core';
import {ListRequestComponent} from './list-request.component';
import {ChangeDetectorRef, CUSTOM_ELEMENTS_SCHEMA} from '@angular/core';
import {SessionService} from "../../../../shared/services/session.service";
import {AppConfigService} from "../../../../services/app-config.service";
import {SearchTriggerService} from "../../../../shared/components/global-search/search-trigger.service";
import {MessageHandlerService} from "../../../../shared/services/message-handler.service";
import {of} from 'rxjs';
import {delay} from 'rxjs/operators';
import {RequestService} from "../../../../shared/services";
import {OperationService} from "../../../../shared/components/operation/operation.service";
import {ConfirmationDialogService} from "../../../global-confirmation-dialog/confirmation-dialog.service";
import {SharedTestingModule} from "../../../../shared/shared.module";

describe('ListRequestComponent', () => {
    let component: ListRequestComponent;
    let fixture: ComponentFixture<ListRequestComponent>;
    const mockRequestService = {
        listRequests: () => {
            return of({
                body: []
            }).pipe(delay(0));
        }
    };
    const mockSessionService = {
        getCurrentUser: () => {
            return false;
        }
    };
    const mockAppConfigService = {
        getConfig: () => {
            return {
                project_creation_restriction: "",
                with_chartmuseum: ""
            };
        }
    };
    const mockSearchTriggerService = {
        closeSearch: () => {
        }
    };
    const mockMessageHandlerService = {
        refresh: () => {
        },
        showSuccess: () => {
        },
    };
    const mockConfirmationDialogService = {
        confirmationConfirm$: of({
            state: "",
            data: []
        })
    };
    const mockOperationService = {
        publishInfo$: () => {
        }
    };
    beforeEach(waitForAsync(() => {
        TestBed.configureTestingModule({
            schemas: [
                CUSTOM_ELEMENTS_SCHEMA
            ],
            imports: [
                SharedTestingModule
            ],
            declarations: [ListRequestComponent],
            providers: [
                TranslateService,
                ChangeDetectorRef,
                {provide: RequestService, useValue: mockRequestService},
                {provide: SessionService, useValue: mockSessionService},
                {provide: AppConfigService, useValue: mockAppConfigService},
                {provide: SearchTriggerService, useValue: mockSearchTriggerService},
                {provide: MessageHandlerService, useValue: mockMessageHandlerService},
                {provide: ConfirmationDialogService, useValue: mockConfirmationDialogService},
                {provide: OperationService, useValue: mockOperationService},
                {provide: ComponentFixtureAutoDetect, useValue: true}

            ]
        })
            .compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(ListRequestComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    let originalTimeout;
    beforeEach(function () {
        originalTimeout = jasmine.DEFAULT_TIMEOUT_INTERVAL;
        jasmine.DEFAULT_TIMEOUT_INTERVAL = 100000;
    });

    afterEach(function () {
        jasmine.DEFAULT_TIMEOUT_INTERVAL = originalTimeout;
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
