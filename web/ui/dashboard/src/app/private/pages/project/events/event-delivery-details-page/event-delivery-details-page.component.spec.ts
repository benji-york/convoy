import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { EventDeliveryDetailsPageComponent } from './event-delivery-details-page.component';

describe('EventDeliveryDetailsPageComponent', () => {
  let component: EventDeliveryDetailsPageComponent;
  let fixture: ComponentFixture<EventDeliveryDetailsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EventDeliveryDetailsPageComponent ],
      imports: [RouterTestingModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EventDeliveryDetailsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
